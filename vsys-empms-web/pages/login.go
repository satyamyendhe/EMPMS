package pages

import (
	u "vsys.empms.commons/utils"
)

type GetLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ErrorMsg string `json:"errorMsg"`
}

func (l *GetLogin) Build() string {
	return u.JoinStr(`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Login</title>
  <!-- Stylesheets -->
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
  <!-- Main Content -->
    <main class="main" id="top">
    <div class="theme-control-toggle fa-icon-wait d-flex justify-content-end">
        <input class="form-check-input ms-0 theme-control-toggle-input" id="themeControlToggle" type="checkbox"
            data-theme-control="theme" value="dark" />
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
                                  <h2>Log in</h2>
                              </div>
                              <div class="col-auto fs--1 text-600">
                                  <span class="mb-0 undefined">or</span>
                                  <span><a href="/signup">Create an account</a></span>
                              </div>
                          </div>
                          <form method="POST" action="/login" >
                              <div class="mb-3">
                                  <input class="form-control" type="email" name="email" placeholder="Email address" id="userEmail" required/>
                              </div>
                              <div class="mb-3">
                                  <input class="form-control" type="password" name="password" placeholder="Password" id="userPass" required/>
                              </div>
                              <div class="errorMsg">`, l.ErrorMsg, `</div>
                              <div class="row flex-between-center">
                                  <div class="col-auto">
                                      <div class="form-check mb-0">
                                          <input class="form-check-input" type="checkbox" id="basic-checkbox" checked="checked" />
                                          <label class="form-check-label mb-0" for="basic-checkbox">Remember me</label>
                                      </div>
                                  </div>
                                  <div class="col-auto"><a class="fs--1" href="">Forgot Password?</a></div>
                              </div>
                              <div class="mb-3">
                                  <button class="btn btn-primary d-block w-100 mt-3" type="submit">Log in</button>
                              </div>
                          </form>
                      </div>
                  </div>
              </div>
          </div>
      </div>
  </main>
  <!-- End of Main Content -->

       
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
