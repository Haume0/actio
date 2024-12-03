package actio

// * Hostname bilgileri

// Hostname is the hostname of the API
var Hostname string = "http://localhost:3000/api"

// Sitename is the site name
var Sitename string = "http://localhost:3000"

// NameRegex is the regex for name
var NameRegex string = `^[a-zA-Z0-9ğüşıöçĞÜŞİÖÇ]{4,64}$`

// PasswordRegex is the regex for password
var PasswordRegex string = `/^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-.]).{8,}$/`

// MailRegex is the regex for mail
var MailRegex string = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

// Languages: Available Languages
var Languages = []string{"tr", "en"}
