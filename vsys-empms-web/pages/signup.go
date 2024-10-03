package pages

import (
	u "vsys.empms.commons/utils"
)

type GetSignUp struct {
	ErrorMsg string `json:"errorMsg"`
}

func (l *GetSignUp) Build() string {
	return u.JoinStr(`
<!DOCTYPE html>
<html lang="en-US" dir="ltr">

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <!-- ===============================================-->
  <!--    Document Title-->
  <!-- ===============================================-->
  <title>Signup</title>

  <!-- ===============================================-->
  <!--    Stylesheets-->
  <!-- ===============================================-->
  <link href="../static/falcon/public/assets/css/theme.min.css" rel="stylesheet" id="style-default">
  <link href="../static/falcon/public/assets/css/user-rtl.min.css" rel="stylesheet" id="user-style-rtl">
  <link href="../static/falcon/public/assets/css/user.min.css" rel="stylesheet" id="user-style-default">

    <style>
      .errorMsg {
          color: red;
      }
      .theme-control-toggle {
          position: absolute;
          top: 10px;
          right: 10px;
          z-index: 1000;
      }
  </style>

  <!-- ===============================================-->
  <!--    Favicons-->
  <!-- ===============================================-->
  
  <!-- <link rel="manifest" href="../../assets/img/favicons/manifest.json"> -->
  <script src="../static/falcon/public/assets/js/config.js"></script>
  <script src="../static/falcon/public/vendors/simplebar/simplebar.min.js"></script>
  
  
  <link rel="stylesheet" href="../static/falcon/public/vendors/fontawesome/all.min.js">

      <!-- BOX ICONS CSS -->
      <link rel="stylesheet" href="../static/boxicons/css/boxicons.css">
  <!-- ===============================================-->
  <!--    Stylesheets-->
  <!-- ===============================================-->

  <link href="../static/falcon/public/assets/css/theme.min.css" rel="stylesheet" id="style-default">
  
</head>

<body>

  <!-- ===============================================-->
  <!--    Main Content-->
  <!-- ===============================================-->
  <main class="main" id="top">

    <div class="theme-control-toggle fa-icon-wait d-flex justify-content-end">
        <input class="form-check-input ms-0 theme-control-toggle-input" id="themeControlToggle" type="checkbox"
            data-theme-control="theme" value="dark"/>
        <label class="mb-0 theme-control-toggle-label theme-control-toggle-light" for="themeControlToggle"
            data-bs-toggle="tooltip" data-bs-placement="left" title="Switch to light theme"><span
        class="fas fa-sun fs-0"></span></label>
        <label class="mb-0 theme-control-toggle-label theme-control-toggle-dark" for="themeControlToggle"
            data-bs-toggle="tooltip" data-bs-placement="left" title="Switch to dark theme"><span
        class="fas fa-moon fs-0"></span></label>
    </div>


    <div class="container" data-layout="container">
      <div class="row flex-center min-vh-100 py-6">
        <div class="col-sm-10 col-md-8 col-lg-6 col-xl-5 col-xxl-4"><a class="d-flex flex-center mb-4"><span class="font-sans-serif fw-bolder fs-5 d-inline-block">EMPMS</span></a>
          <div class="card">
            <div class="card-body p-4 p-sm-5">
              <div class="row flex-between-center mb-2">
                <div class="col-auto">
                  <h2>Register</h2>
                </div>
                <div class="col-auto fs--1 text-600"><span class="mb-0 undefined">Have an account?</span> <span><a href="/login">Login</a></span></div>
              </div>
              <form id="registrationForm" method="POST" action="/signup">
                <div class="mb-3">
                  <input class="form-control" type="text" name="name" id="name" autocomplete="on" placeholder="Name" required/>
                </div>
                <div class="mb-3">
                  <input class="form-control" type="email" name="email" id="email" autocomplete="on" placeholder="Email address" required/>
                </div>
                <div class="errorMsg">`, l.ErrorMsg, `</div>
                <div class="row gx-2">
                  <div class="mb-3 col-sm-6">
                    <input class="form-control" type="password" name="password" id="password" autocomplete="on" placeholder="Password" required/>
                  </div>
                  <div class="mb-3 col-sm-6">
                    <input class="form-control" type="password" id="confirmPassword" autocomplete="on" placeholder="Confirm Password" required/>
                  </div>
                </div>
                <div class="form-check">
                  <input class="form-check-input" type="checkbox" id="basic-register-checkbox" />
                  <label class="form-label" for="basic-register-checkbox">I accept the <a href="#!">terms </a>and <a href="#!">privacy policy</a></label>
                </div>
                <div class="mb-3">
                  <button class="btn btn-primary d-block w-100 mt-3" type="submit">Register</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
  <!-- ===============================================-->
  <!--    End of Main Content-->
  <!-- ===============================================-->

  <script>
    document.addEventListener("DOMContentLoaded", () => {
      document.getElementById('registrationForm').addEventListener('submit', function(event) {
        event.preventDefault(); 
        
        var passwordField = document.getElementById('password');
        var confirmPasswordField = document.getElementById('confirmPassword');

        var password = passwordField.value;
        var confirmPassword = confirmPasswordField.value;

        // Check if the passwords match
        if (password !== confirmPassword) {
          // Alert and focus on the password fields
          alert('Passwords do not match.');
          passwordField.value = "";
          confirmPasswordField.value = "";
          passwordField.focus();
        } else {
          // Submit the form if passwords match
          this.submit();
        }
      });
    });
  </script>
    
  
    <script src="../static/falcon/public/vendors/popper/popper.min.js"></script>
    <script src="../static/falcon/public/vendors/bootstrap/bootstrap.min.js"></script>
    <script src="../static/falcon/public/vendors/anchorjs/anchor.min.js"></script>
    <script src="../static/falcon/public/vendors/is/is.min.js"></script>
    <script src="../static/falcon/public/vendors/prism/prism.js"></script>
    <script src="../static/falcon/public/vendors/fontawesome/all.min.js"></script>
    <script src="../static/falcon/public/vendors/lodash/lodash.min.js"></script>
    <script src="../static/falcon/public/vendors/list.js/list.min.js"></script>
    <script src="../static/falcon/public/assets/js/theme.js"></script>

</body>

</html>
	`)
}
