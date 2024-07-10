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
		pgxUseCase.CommentUseCase
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
	userUseCase := pgxUseCase.NewUserUseCase(userRepo, postRepo, commentRepo, transactor)
	postUseCase := pgxUseCase.NewPostUseCase(postRepo, commentRepo, transactor)
	commentUseCase := pgxUseCase.NewCommentUseCase(commentRepo, postRepo, userRepo, transactor)

	return &Interactor{
		UserRepository:    userRepo,
		PostRepository:    postRepo,
		CommentRepository: commentRepo,
		UserUseCase:       userUseCase,
		PostUseCase:       postUseCase,
		CommentUseCase:    commentUseCase,
	}
}
