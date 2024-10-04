package models

type Product struct {
    ID			string	`json:"id"`
    Name  		string	`json:"name"`
    Price		int 	`json:"Price"`
    Category	string	`json:"category"`
    Brand		string	`json:"brand"`
    Rating		int		`json:"rating"`
    Selected	bool	`json:"selected"`
    Ordered		bool	`json:"ordered"`
}


// type Product struct {
//     ID			int64	`json:"id"`
//     Name  		string	`json:"name"`
//     Price		float32 `json:"Price"`
//     Category	string	`json:"category"`
//     Brand		string	`json:"brand"`
//     Rating		float32	`json:"rating"`
//     Selected	bool	`json:"selected"`
//     Ordered		bool	`json:"ordered"`
// }



