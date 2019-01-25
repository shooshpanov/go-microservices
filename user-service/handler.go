package main

import  (
	"golang.org/x/net/context"
	pb "github.com/shooshpanov/microservices-project/user-service/proto/user"
)

type service struct {
	repo Repository
	tokenService Authable
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.id)
	if err != nil {
		returm err
	}
	res.User = user
	return nil
}

func (srv *serivce) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	user, err := srv.repo.GetEmailAndPassword(req)
	if err != nil {
		return nil
	}
	res.Token = "testingabs"
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err = srv.repo.Create(req); err != nil {
		return nil
	}
	res.User = req
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}