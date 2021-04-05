<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .Pagetitle }} | Open Function Server</title>
    <base href="http://localhost:9099{{ .RoutePrefix }}/">

    <!-- TODO: build static assets locally and set up route to serve them -->
    <link rel="icon" href="/images/favicon.svg">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/5.0.0-beta1/css/bootstrap.min.css" integrity="sha512-thoh2veB35ojlAhyYZC0eaztTAUhxLvSZlWrNtlV01njqs/UdY3421Jg7lX0Gq9SRdGVQeL8xeBp9x1IPyL1wQ==" crossorigin="anonymous" />

    <script src="https://cdnjs.cloudflare.com/ajax/libs/alpinejs/2.8.0/alpine-ie11.min.js" integrity="sha512-ao6Tt3nuCMC5fwlPf9guyKVyQn8vd82mOIMSsiTManuRb8tYqTocCglR6nk2NLK4+uzMOem+hTLuobarbn/RMg==" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.9.4/Chart.min.js" integrity="sha512-d9xgZrVZpmmQlfonhQUvTR7lMPtO7NkZMkA0ABN3PHCbKA5nqylQ/yWlFAyY6hYgdF1Qh6nYiuADWwKB4C2WSw==" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/5.0.0-beta1/js/bootstrap.min.js" integrity="sha512-ZvbjbJnytX9Sa03/AcbP/nh9K95tar4R0IAxTS2gh2ChiInefr1r7EpnVClpTWUEN7VarvEsH3quvkY1h0dAFg==" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/axios/0.21.1/axios.min.js" integrity="sha512-bZS47S7sPOxkjU/4Bt0zrhEtWx0y0CRkhEp8IckzK+ltifIIE9EMIMTuT/mEzoIMewUINruDBIR/jJnbguonqQ==" crossorigin="anonymous"></script>
</head>
<body class="{{ .BodyClasses }}">

<div class="container-fluid">
    <div class="row" style="min-height: 100vh;">
        <div class="col-2 sidebar" style="background-color: black; color: white;">
            <p class="pt-4"><strong>Open Function Server Maintenance</strong></p>
            <nav>
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <a class="nav-link border-left" href="#">Dashboard</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link border-left" href="sites">Sites</a>
                    </li>
                </ul>
            </nav>
        </div>
        <div class="col-10 main-content">
