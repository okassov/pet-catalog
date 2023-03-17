package usecase

type (
	// Auth interface {
	// 	SignUp(context.Context, entity.User) error
	// 	SignIn(context.Context, entity.User) (map[string]string, error)
	// 	ValidateToken(ctx context.Context, accessToken string) (UserClaim, error)
	// }

	Catalog interface {
	}
	// Product interface {
	// }

	// Category interface {
	// }

	ProductRepo interface {
		// CreateUser(context.Context, entity.User) error
		// GetUser(context.Context, entity.User) (*entity.User, error)
	}

	CategoryRepo interface {
		// CreateUser(context.Context, entity.User) error
		// GetUser(context.Context, entity.User) (*entity.User, error)
	}
)
