# go-recaptcha
Golang client library for reCAPTCHA. https://www.google.com/recaptcha/intro/v3.html

# HOWTO

1. get package

   ```bash
   $ go get -u github.com/czsilence/go-recaptcha/recaptcha
   ```

2. use in project

   ```golang

    import "github.com/czsilence/go-recaptcha/recaptcha"

    // init config
    func some_init_func_before_verify() {
        // other code
        recaptcha.Config("your secret")
        // other code
    }

    // verify
    func verifyfunc(token string) {
       // other code

       res, err := recaptcha.Verify(token)
       // process res and err

       // other code
    }

   ```