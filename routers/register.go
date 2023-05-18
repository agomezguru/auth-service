package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agomezguru/cti-b2b/auth-service/db"
	"github.com/agomezguru/cti-b2b/auth-service/models"
	"github.com/thedevsaddam/govalidator"
)

/* User register routine */
func Register(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")
	var newUser models.User 

	// Start validations
	rules := govalidator.MapData{
		"login": 		[]string{"required", "between:4,30", "alpha_dash"},
		"name": 		[]string{"required", "min:2", "max:255", "regex:^[a-zA-ZñÑáéíóúÁÉÍÓÚ ]+$"},
		"surname":	[]string{"required", "min:2", "max:255", "regex:^[a-zA-ZñÑáéíóúÁÉÍÓÚ ]+$"},
		"email":    []string{"required", "min:4", "max:100", "email"},
		"password":	[]string{"required", "min:12", "max:100", "regex:^[0-9a-zA-ZñÑáéíóúÁÉÍÓÚ #?!@$%^&.*-_]+$"},
	}

	/*
	$rules = [
		'password' => 'required|string|min:12|confirmed|regex:/^[0-9a-zA-ZñÑáéíóúÁÉÍÓÚ #?!@$%^&.*-_]+$/i',
		'rol'      => 'required',
	]; */

	messages := govalidator.MapData{
		"login": 		[]string{"required:You must provide login", "between:The username field must be between 4 to 30 chars"},
		"name":			[]string{"regex:Allowed only letters in name"},
		"surname":	[]string{"regex:Allowed only letters in surname"},
		"password":	[]string{"required:You mut provide password", "regex:Password must include at least: 1 uppercase, 1 lowercase, 1 number and 1 character special # ? ! @ $ % ^ & . * - _"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Data:						 &newUser,
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON() // Incoming JSON data in Go data struct
	message := map[string]interface{}{"validationError": e}

	// Incoming JSON data with errors
	if fmt.Sprint(e) !=  "map[]" {
		w.WriteHeader(http.StatusBadRequest)
		
		json.NewEncoder(w).Encode(message)
		return
	}

	_, finded, _ := db.UserExist(newUser.Login, newUser.Email)
	if finded == true {
		errDescription := "One user was previously registered with this login or email."
		errMap := map[string]string {"email": errDescription}
		message = map[string]interface{}{"validationError": errMap}
		w.WriteHeader(http.StatusBadRequest)
		
		json.NewEncoder(w).Encode(message)
		return
	}

	// All validations passed. Try to register a new user.
	status, err := db.InsertRegister(newUser)
	if err != nil {
		errDescription := "Error ocurred when try to register a new user: " + err.Error()
		errMap := map[string]string {"response": errDescription}
		message = map[string]interface{}{"databaseError": errMap}
		w.WriteHeader(http.StatusBadRequest)
		
		json.NewEncoder(w).Encode(message)
		return
	}

	if status == false {
		errDescription := "Insert register in DB failed."
		errMap := map[string]string {"response": errDescription}
		message = map[string]interface{}{"databaseError": errMap}
		w.WriteHeader(http.StatusBadRequest)
		
		json.NewEncoder(w).Encode(message)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
