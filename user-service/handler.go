package main

import  (
	"golang.org/x/net/context"
	pb "github.com/shooshpanov/microservices-project/user-service/proto/user"
)

const topic = "user.created"

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
	log.Println("Logging in with:", req.Email, req.Password)
	user, err := srv.repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	res.User = req

	if err := srv.Publisher.Publish(ctx, req); err != nil {
		return err
	}

	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

	// Decode token
	claims, err := srv.tokenService.Decode(req.Token)

	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}

// func (srv *servuce) publishEvent(user *pb.User) error {
// 	// Marshall JSON string
// 	body, err := json.Marshal(user)
// 	if err != nil {
// 		return err
// 	}

// 	// Create message to broker
// 	msg := &broker.Message{
// 		Header: map[string]string{
// 			"id": user.id,
// 		},
// 		Body: body,
// 	}

// 	// Publish message to broker
// 	if err:= srv.PubSub.Publish(topic, msg); err != nil {
// 		log.Printf("[pub] failed: %v", err)
// 	}

// 	return nil
// }