package dto

import (
	z "github.com/Oudwins/zog"
)

var SignupSchema = z.Struct(
	z.Shape{
		"name":     z.String().Min(3, z.Message("Name must contain minimum 3 characters")).Max(10, z.Message("Name must contain maximum 10 characters")).Required(),
		"email":    z.String().Email().Required(z.Message("Email is required")),
		"password": z.String().Min(8, z.Message("Password should be of atleast 8 characaters")).Max(16, z.Message("Password should be of atmost 16 characters")).Required(),
		"fullname": z.String(),
	},
)

var WebSiteSchema = z.Struct(
	z.Shape{
		"url": z.String().URL().Required(),
	},
)
