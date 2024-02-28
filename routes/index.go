package routes

import (
	"LibraryManagement/controller"
	"LibraryManagement/database"
	"LibraryManagement/middleware"
	"LibraryManagement/model"
	"LibraryManagement/repository"
	"LibraryManagement/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"net/http"
	"strconv"
)

var (
	db = database.Connect()

	userRepository     = repository.UserRepository(db)
	bookRepository     = repository.NewBookRepository(db)
	addressRepository  = repository.NewAddressRepository(db)
	libraryRepository  = repository.NewLibraryRepository(db)
	checkoutRepository = repository.NewCheckoutRepository(db)

	authService    = service.AuthService(userRepository)
	jwtService     = service.NewJwtService()
	userService    = service.NewUserService(userRepository, checkoutRepository)
	adminService   = service.NewAdminService(bookRepository)
	bookService    = service.NewBookService(bookRepository)
	addressService = service.NewAddressService(addressRepository)
	libraryService = service.NewLibraryService(libraryRepository)

	authController    = controller.AuthController(authService, jwtService)
	userController    = controller.NewUserController(userService)
	adminController   = controller.NewAdminController(bookService)
	generalController = controller.NewGeneralController(addressService, libraryService)
)

func AddRoutes(superRoute *gin.RouterGroup) {
	authRoutes := superRoute.Group("/auth")
	AuthRoutes(authRoutes, authController)

	userRoutes := superRoute.Group("/user")
	userRoutes.Use(middleware.AuthorizeJWT(jwtService))
	UserRoutes(userRoutes, userController)

	adminRoutes := superRoute.Group("/admin")
	adminRoutes.Use(middleware.AuthorizeJWT(jwtService))
	AdminRoutes(adminRoutes, adminController)

	superRoute.GET("/libraries/:address_name", generalController.GetLibrariesOfAddress)
	superRoute.GET("/books/:library_id", generalController.GetBooksOfLibrary)
	superRoute.GET("/get-pdf-report", func(context *gin.Context) {
		checkouts, err := checkoutRepository.CountCheckoutByCurrentYearGroupByMonth()
		if err != nil {
			response := model.BuildResponse("Failed to process request", err.Error(), nil)
			context.AbortWithStatusJSON(400, response)
			return
		}
		generatePdfFile(checkouts)
		context.FileAttachment("reports/report.pdf", "report.pdf")
	})
	superRoute.GET("/send-an-email", func(context *gin.Context) {
		sender := service.NewSender()
		message := service.NewMessage("Hello", "This is a test email")
		message.To = []string{"danghuutan264@gmail.com"}
		err := message.AttachFile("reports/report.pdf")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		fmt.Println(sender.Send(message))
	})
}

func generatePdfFile(data map[string][2]uint) error {
	//var _ float64 = 40
	//var height float64 = 10
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// write title
	pdf.SetFont("Arial", "B", 16)
	pageWidth, _ := pdf.GetPageSize()
	titleText := "Report"
	titleWidth := pdf.GetStringWidth(titleText)
	centerX := (pageWidth - titleWidth) / 2
	pdf.SetX(centerX)
	pdf.Write(25.4, titleText)

	// write column headers
	pdf.SetFont("Arial", "", 14)
	colWidths := []float64{25, 70, 70, 25}
	colNames := []string{"Month", "Borrowed Quantity", "Returned Quantity", "Rate"}
	pdf.SetX(25.4)
	pdf.SetY(pdf.GetY() + 25)
	pdf.SetFillColor(255, 255, 255)
	for i, width := range colWidths {
		pdf.Rect(pdf.GetX(), pdf.GetY(), width, 20, "FD")
		pdf.CellFormat(width, 20, colNames[i], "1", 0, "C", true, 0, "")
	}
	pdf.SetY(pdf.GetY() + 20)

	// write data
	pdf.SetFont("Arial", "", 12)
	for key, value := range data {
		for i, width := range colWidths {
			content := func() string {
				if i == 0 {
					return key
				} else if i == 1 {
					return strconv.Itoa(int(value[0]))
				} else if i == 2 {
					return strconv.Itoa(int(value[1]))
				} else {
					if value[1] == 0 {
						return "none"
					}
					rate := float64(value[0]) / float64(value[1])
					return strconv.FormatFloat(rate, 'f', 2, 64)
				}
			}
			pdf.Rect(pdf.GetX(), pdf.GetY(), width, 20, "FD")
			pdf.CellFormat(width, 20, content(), "1", 0, "C", true, 0, "")
		}
		pdf.SetY(pdf.GetY() + 20)
	}

	err := pdf.OutputFileAndClose("reports/report.pdf")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
