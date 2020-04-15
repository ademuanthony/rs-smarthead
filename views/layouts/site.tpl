    <!DOCTYPE html>
    <html lang="en">
    <head>
        <title>
            Remote School - {{block "title" .}}{{end}}
        </title>

        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <meta name="description" content="{{block "description" .}}{{end}} ">
        <meta name="author" content="Merry World - Ademu Anthony">

        <link rel="icon" type="image/png" sizes="16x16" href="/static/images/remote-school-icon.png">

        <!-- ============================================================== -->
        <!-- Custom fonts for this template                                 -->
        <!-- ============================================================== -->
        <script src="https://kit.fontawesome.com/670ea91c67.js"></script>
        <link href="https://fonts.googleapis.com/css?family=Nunito:200,200i,300,300i,400,400i,600,600i,700,700i,800,800i,900,900i" rel="stylesheet">

        <!-- Global site tag (gtag.js) - Google Analytics -->
        <script async src="https://www.googletagmanager.com/gtag/js?id=UA-158909227-2"></script>
        <script>
        window.dataLayer = window.dataLayer || [];
        function gtag(){dataLayer.push(arguments);}
        gtag('js', new Date());

        gtag('config', 'UA-158909227-2');
        </script>


        <link rel=stylesheet href=https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css integrity=sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm crossorigin=anonymous>
        <link href="//maxcdn.bootstrapcdn.com/font-awesome/4.2.0/css/font-awesome.min.css" rel="stylesheet">
        <!-- ============================================================== -->
        <!-- Base styles for Start Bootstrap template SB Admin 2            -->
        <!-- ============================================================== -->
        <link href="/static/assets/css/sb-admin-2.css" rel="stylesheet">

        <!-- ============================================================== -->
        <!-- Custom styles for this service applied to all pages            -->
        <!-- ============================================================== -->
        <link href="/static/assets/css/custom.css" id="theme" rel="stylesheet">

        <!-- ============================================================== -->
        <!-- Page specific CSS                                              -->
        <!-- ============================================================== -->
        {{block "style" .}} {{end}}

        <!-- Facebook Pixel Code -->
        <script>
        !function(f,b,e,v,n,t,s)
        {if(f.fbq)return;n=f.fbq=function(){n.callMethod?
        n.callMethod.apply(n,arguments):n.queue.push(arguments)};
        if(!f._fbq)f._fbq=n;n.push=n;n.loaded=!0;n.version='2.0';
        n.queue=[];t=b.createElement(e);t.async=!0;
        t.src=v;s=b.getElementsByTagName(e)[0];
        s.parentNode.insertBefore(t,s)}(window, document,'script',
        'https://connect.facebook.net/en_US/fbevents.js');
        fbq('init', '806165446232113');
        fbq('track', 'PageView');
        </script>
        <noscript><img height="1" width="1" style="display:none"
        src="https://www.facebook.com/tr?id=806165446232113&ev=PageView&noscript=1"
        /></noscript>
        <!-- End Facebook Pixel Code -->
    </head>
    <body id="page-top">
    <!-- ============================================================== -->
    <!-- Topbar                                                         -->
    <!-- ============================================================== -->
         <!-- Topbar -->
    
         <div class="fixed-top">
            <header class="topbar">
                <div class="container">
                    <div class="row">
                        <!-- social icon-->
                        <div class="col-sm-12">
                            <ul class="social-network">
                                <li><a class="waves-effect waves-dark" href="#"><i class="fa fa-facebook"></i></a></li>
                                <li><a class="waves-effect waves-dark" href="#"><i class="fa fa-twitter"></i></a></li>
                                <li><a class="waves-effect waves-dark" href="#"><i class="fa fa-linkedin"></i></a></li>
                                <li><a class="waves-effect waves-dark" href="#"><i class="fa fa-instagram"></i></a></li>
                            </ul>
                        </div>
                    </div>
                </div>
            </header>
            <nav class="navbar navbar-expand-lg navbar-dark mx-background-top-linear">
                <div class="container">
                <a class="navbar-brand" href="/" style="text-transform: uppercase;"> 
                    <img alt="REMOTE SCHOOL" class="logo" style="height:40px;" 
                        src="/static/images/remote-school-logo.png">
                </a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarResponsive">
    
                    <ul class="navbar-nav ml-auto">
    
                    <li class="nav-item {{block "homeActive" .}}{{end}}">
                        <a class="nav-link" href="/">Home
                        <span class="sr-only">(current)</span>
                        </a>
                    </li>
    
                    <li class="nav-item {{block "pricingActive" .}}{{end}}">
                        <a class="nav-link" href="/pricing">Pricing</a>
                    </li>
    
                    <li class="nav-item {{block "supportActive" .}}{{end}}">
                        <a class="nav-link" href="/support">Support</a>
                    </li>
    
                    <li class="nav-item">
                        <a class="nav-link" href="https://forms.gle/hnXFyFMem68JJujj8">Register</a>
                    </li>
    
                    </ul>
                </div>
                </div>
            </nav>
        </div>
    <!-- End of Topbar -->
    <!-- End of Topbar -->

        <!-- ============================================================== -->
        <!-- Page Wrapper                                                   -->
        <!-- ============================================================== -->
        <div id="wrapper" class="website">

            <!-- ============================================================== -->
            <!-- Content Wrapper                                                -->
            <!-- ============================================================== -->
            <div id="content-wrapper" class="d-flex flex-column bg-white">

                <!-- ============================================================== -->
                <!-- Main Content                                                   -->
                <!-- ============================================================== -->
                <div id="content">



                    <!-- ============================================================== -->
                    <!-- Page Content                                                   -->
                    <!-- ============================================================== -->


                        {{ template "content" . }}

                    <!-- End Page Content  -->

                </div>
                <!-- End of Main Content -->

                <!-- ============================================================== -->
                <!-- Footer                                                         -->
                <!-- ============================================================== -->
                  <!-- Site footer -->
  <footer class="site-footer">
    <div class="container">
      <div class="row">
        <div class="col-sm-12 col-md-6">
          <h6>About</h6>
          <p class="text-justify">Remoteschool.com.ng
               is a cloud-based learning platform that offers safe and affordable 
               interactive online teaching and learning</p>
        </div>

        <div class="col-xs-6 col-md-3"> 
          <h6>Top Subject</h6> 
          <ul class="footer-links">
            <li><a href="/signup">Maths</a></li>
            <li><a href="/signup">English</a></li>
            <li><a href="/signup">Biology</a></li>
            <li><a href="/signup">Private classes</a></li>
          </ul>
        </div>

        <div class="col-xs-6 col-md-3">
          <h6>Quick Links</h6> 
          <ul class="footer-links">
            <li><a class="p-2 text-white" href="/">Home</a></li>
            <li><a class="p-2 text-white" href="/pricing">Pricing</a></li>
            <li><a class="p-2 text-white" href="/support">Support</a></li>
            <li><a class="p-2 text-white" href="https://forms.gle/xgnAREA3fByZYKuc6" target="_blank">Apply as Teacher</a></li>
          </ul>
        </div>
      </div>
      <hr>
    </div>
    <div class="container">
      <div class="row">
        <div class="col-md-8 col-sm-6 col-xs-12">
          <p class="copyright-text">Copyright &copy; 2020 All Rights Reserved by 
       <a href="#">Merry World (BN 2662048)</a>.
          </p>
        </div>

        <div class="col-md-4 col-sm-6 col-xs-12">
          <ul class="social-icons">
            <li><a class="facebook" href="#"><i class="fa fa-facebook"></i></a></li>
            <li><a class="twitter" href="#"><i class="fa fa-twitter"></i></a></li>
            <li><a class="dribbble" href="#"><i class="fa fa-instagram"></i></a></li>
            <li><a class="linkedin" href="#"><i class="fa fa-linkedin"></i></a></li>   
          </ul>
        </div>
      </div>
    </div>
</footer>
<!--Start of Tawk.to Script-->
<script type="text/javascript">
var Tawk_API=Tawk_API||{}, Tawk_LoadStart=new Date();
(function(){
var s1=document.createElement("script"),s0=document.getElementsByTagName("script")[0];
s1.async=true;
s1.src='https://embed.tawk.to/5e83343d35bcbb0c9aac3887/default';
s1.charset='UTF-8';
s1.setAttribute('crossorigin','*');
s0.parentNode.insertBefore(s1,s0);
})();
</script>
<!--End of Tawk.to Script-->
                <!-- End of Footer -->

            </div>
            <!-- End of Content Wrapper -->

        </div>
        <!-- End of Page Wrapper -->

        <!-- Scroll to Top Button-->
        <a class="scroll-to-top rounded" href="#page-top">
            <i class="fas fa-angle-up"></i>
        </a>


    <!-- ============================================================== -->
    <!-- Javascript Bootstrap core JavaScript                           -->
    <!-- ============================================================== -->
    <script src="/static/assets/vendor/jquery/jquery.min.js"></script>
    <script src="/static/assets/vendor/bootstrap/js/bootstrap.bundle.min.js"></script>

    <!-- ============================================================== -->
    <!-- Core plugin JavaScript                                         -->
    <!-- ============================================================== -->
    <script src="/static/assets/vendor/jquery-easing/jquery.easing.min.js"></script>

    <!-- ============================================================== -->
    <!-- Javascript for Start Bootstrap template SB Admin 2             -->
    <!-- ============================================================== -->
    <script src="/static/assets/js/sb-admin-2.min.js"></script>

    <!-- ============================================================== -->
    <!-- Page specific Javascript                                       -->
    <!-- ============================================================== -->
    {{block "js" .}} {{end}}
    </body>
    </html>