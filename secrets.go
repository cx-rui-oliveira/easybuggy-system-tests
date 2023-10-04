package easybuggy4sb

import os

// Sample secrets
api_key = "sk_test_0123456789abcdef0123456789abcdef"
password = "mySuperSecretPassword"
database_url = "postgresql://username:password@localhost/mydatabase"
aws_access_key = "AKIA1234567890EXAMPLE"
aws_secret_key = "abc12345def67890ghiJKLmnoPQRSTuVWxyz"
encryption_key = "e1f8abc23e1f8abc23e1f8abc23e1f8a"
auth_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIn0.lw2-MlExE-xBKYk8Hhgyb28OR_vGZrT73YKPfYx8QBE"

def get_api_key():
return "sk_test_0123456789abcdef0123456789abcdef"

def get_password():
return "mySuperSecretPassword"

def connect_to_database():
return "postgresql://username:password@localhost/mydatabase"

// AWS credentials
os.environ["AWS_ACCESS_KEY_ID"] = "AKIA1234567890EXAMPLE"
os.environ["AWS_SECRET_ACCESS_KEY"] = "abc12345def67890ghiJKLmnoPQRSTuVWxyz"

// Encryption key
encryption_key_variable = "e1f8abc23e1f8abc23e1f8abc23e1f8a"

// JWT authentication token
auth_token_variable = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIn0.lw2-MlExE-xBKYk8Hhgyb28OR_vGZrT73YKPfYx8QBE"
