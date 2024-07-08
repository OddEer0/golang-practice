package pgxInteractor

import (
	pgxPractice "github.com/OddEer0/golang-practice/practices/pgx_practice"
	pgxRepository "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_repository"
	pgxUseCase "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_usecase"
	"github.com/OddEer0/golang-practice/resources/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Interactor struct {
		UserRepository    repository.User
		PostRepository    repository.Post
		CommentRepository repository.Comment
		pgxUseCase.UserUseCase
		pgxUseCase.PostUseCase
	}
)

func New(db *pgxpool.Pool) *Interactor {
	txController := pgxPractice.PgxTransactionController{}
	transactor := &pgxPractice.PgxTransaction{
		Conn:         db,
		TxController: txController,
	}
	userRepo := pgxRepository.NewUserRepository(db, txController)
	postRepo := pgxRepository.NewPostRepository(db, txController)
	commentRepo := pgxRepository.NewCommentRepository(db, txController)
	userUseCase := pgxUseCase.NewUserUseCase(userRepo, postRepo, transactor)
	postUseCase := pgxUseCase.NewPostUseCase(postRepo, commentRepo, transactor)

	return &Interactor{
		UserRepository: userRepo,
		PostRepository: postRepo,
		UserUseCase:    userUseCase,
		PostUseCase:    postUseCase,
	}
}
