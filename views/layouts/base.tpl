<!DOCTYPE html>
    <html lang="en">
    <head>
        <title>
            {{block "title" .}}{{end}} - Remote School
        </title>

        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <meta name="description" content="Remote online interactive teaching and learning platform.">
        <meta name="author" content="Merry World Media">

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


        <!-- ============================================================== -->
        <!-- Base styles for Start Bootstrap template SB Admin 2            -->
        <!-- ============================================================== -->
        <link href="/static/assets/css/sb-admin-2.css" rel="stylesheet">

        <!-- ============================================================== -->
        <!-- Custom styles for this service applied to all pages            -->
        <!-- ============================================================== -->
        <link href="/static/assets/css/custom.css" rel="stylesheet">

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

        {{ template "partials/app-wrapper" . }}

        <!-- ============================================================== -->
        <!--  Logout Modal                                                  -->
        <!-- ============================================================== -->
        {{ if eq 1 1}}
        <div class="modal fade" id="logoutModal" tabindex="-1" role="dialog" aria-labelledby="logoutModalLabel" aria-hidden="true">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title" id="logoutModalLabel">Ready to Leave?</h5>
                        <button class="close" type="button" data-dismiss="modal" aria-label="Close">
                            <span aria-hidden="true">×</span>
                        </button>
                    </div>
                    <div class="modal-body">Select "Logout" below if you are ready to end your current session.</div>
                    <div class="modal-footer">
                        <button class="btn btn-secondary" type="button" data-dismiss="modal">Cancel</button>
                        <a class="btn btn-primary" href="/user/logout">Logout</a>
                    </div>
                </div>
            </div>
        </div>
        {{ end }} 
 
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
        <script src="/static/assets/js/sb-admin-2.js"></script>

        <!-- ============================================================== -->
        <!-- Custom Javascript for this service applied to all pages        -->
        <!-- ============================================================== -->
        <script src="/static/assets/js/custom.js"></script>

        <!-- ============================================================== -->
        <!-- Page specific Javascript                                       -->
        <!-- ============================================================== -->
        {{block "js" .}} {{end}}

        <script src="/static/assets/js/vendors~app.bundle.js"></script>
        <script src="/static/assets/js/app.bundle.js="></script>
    </body>
</html>
{{ define "invalid-feedback" }}
   
{{ end }}
{{ define "app-flashes" }}
    
{{ end }}
{{ define "validation-error" }}
    
{{ end }}