package services

import (
	"context"
	"user-panel/internal/db"
	pb "user-panel/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	pb.UnimplementedUserPanelServiceServer
}

// GetUserActivities fetches user activities for gRPC
func (s *GRPCServer) GetUserActivities(ctx context.Context, req *pb.UserRequest) (*pb.UserActivitiesResponse, error) {
	rows, err := db.DB.Query("SELECT activity, created_at FROM user_activities WHERE user_id=$1 ORDER BY created_at DESC", req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []*pb.Activity
	for rows.Next() {
		var activity, createdAt string
		if err := rows.Scan(&activity, &createdAt); err != nil {
			return nil, err
		}
		activities = append(activities, &pb.Activity{
			Activity:  activity,
			CreatedAt: createdAt,
		})
	}

	return &pb.UserActivitiesResponse{Activities: activities}, nil
}

// GetWalletTransactions fetches wallet transactions for gRPC
func (s *GRPCServer) GetWalletTransactions(ctx context.Context, req *pb.UserRequest) (*pb.WalletTransactionsResponse, error) {
	rows, err := db.DB.Query("SELECT transaction_type, amount, description, created_at FROM wallet_transactions WHERE user_id=$1 ORDER BY created_at DESC", req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*pb.Transaction
	for rows.Next() {
		var transactionType, description, createdAt string
		var amount float64
		if err := rows.Scan(&transactionType, &amount, &description, &createdAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, &pb.Transaction{
			TransactionType: transactionType,
			Amount:          amount,
			Description:     description,
			CreatedAt:       createdAt,
		})
	}

	return &pb.WalletTransactionsResponse{Transactions: transactions}, nil
}
