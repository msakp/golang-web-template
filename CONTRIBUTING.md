# CONTRIBUTING

## Error handling

service methods, than communicate or could communicate with handlers should return \*dto.HttpErr instead of plain error  
moreover they should return Error wrappers **like so:**
```go
func (s *someService) Dosmth(ctx context.Context, data any) (newData any, err error){
	if !s.someRepo.Exists(data){
		return nil, wrapper.NotFoundErr(dto.MsgSomeDataNotFound)
	}
	if err != nil{
		return nil, wrapper.InternalServerErr(err.Error())
	}
}
```
note to use common static error messages (like MsgSomeDataNotFound) and populate them( **domain/dto/error.go** ), if needed to remain errors consistency

### **Error Handling from Handlers level**
```go
func (sh *someHandler) Dosmth(c *fiber.Ctx) error{
	var Data any
	if err := c.BodyParser(&Data); err != nil{
		httpErr := wrapper.BadRequestErr(err.Error())
		return c.Status(httpErr.HttpCode).JSON(httpErr)
	}
	data, err := sh.someService.Dosmth(c.UserContext(), Data)
	if err != nil{
		return c.Status(err.HttpCode).JSON(err)
	}
}
```


## Service-Repo communications
- interfaces should use dto to communicate, not storage models (for exmpl storage.User), use wrappers inside repo to convert shit for database.
- repos should return original plain errors, could log them explicitly  

**User Contracts Example**
```go
type UserRepository interface {
	Exists(ctx context.Context, email string) bool
	Create(ctx context.Context, u *dto.UserRegister) (uuid.UUID, error)
	GetByEmail(ctx context.Context, email string) (*dto.UserView, error)
	GetById(ctx context.Context, id uuid.UUID) (*dto.UserView, error)
}

type UserService interface {
	Register(ctx context.Context, u *dto.UserRegister) (*dto.UserAuthResponse, *dto.HttpErr)
	Login(ctx context.Context, uLogin *dto.UserLogin) (*dto.UserAuthResponse, *dto.HttpErr)
	GetProfile(ctx context.Context, email string) (*dto.UserView, *dto.HttpErr)
}
```
note than repo takes **\*dto.UserRepister** and returns **\*dto.UserView**, instead of \*storage.CreateUserParams and \*storage.User


## Logging

try to prefer to log outside of handlers layer(if possible)  

**logger usage example**
```go
func (s *someService) Dosmth(ctx context.Context, data any) (newData any, err *dto.HttpErr){
	if err != nil{
		logger.FromCtx(ctx).Error(ctx, fmt.Sprintf("someService failed to Dosmth with ERR: %s", err.Error()))
		return nil, wrapper.InternalServerErr(err.Error())
	}
}
```

## ValidatorService
working on top of go-playground validator, with idea of making custom validation fields, and using existsed  

**Contract**
```go
type ValidatorService interface {
	ValidateRequestData(dto any) *dto.HttpErr
	ValidateRequestSlice(dtos any) *dto.HttpErr
}
```
note that **dto must be a pointer, dtos must be a slice**  

**Usage Example**
```go
var Data any
if err := c.BodyParser(&Data); err != nil{
	httpErr := wrapper.BadRequestErr(err.Error())
	return c.Status(httpErr.HttpCode).JSON(httpErr)
}
if err := sh.validatorService.ValidateRequestData(&Data); err != nil{
	return c.Status(err.HttpCode).JSON(err)
}
```
( see more datails in **service/validator.go** )



## Overall
**see more details in user service implementation, dto impl, bla bla bla**

