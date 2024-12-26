
        package controllers

        
        type TicketController struct{}

        
        func (tc *TicketController) BuyTicket() string {
            return "Ticket purchased"
        }

        
        func (tc *TicketController) ReturnTicket() string {
            return "Ticket returned"
        }
    