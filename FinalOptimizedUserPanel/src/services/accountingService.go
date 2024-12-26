package main


        import (
            "log"
        )

        

        package services

        func LogActivity(userID string, action string) {
            log.Printf("User %s performed action: %s", userID, action)
        }
    