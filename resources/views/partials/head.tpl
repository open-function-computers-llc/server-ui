<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .Pagetitle }} | Open Function Server</title>
    <base href="http://localhost:9099/tlioqwtisjdiauegliavaw4gjh/">

    <!-- TODO: build static assets locally and set up route to serve them -->
    <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.0/css/bootstrap.min.css">
</head>
<body class="{{ .BodyClasses }}">

<div class="container-fluid">
    <div class="row" style="min-height: 100vh;">
        <div class="col-2 sidebar" style="background-color: black; color: white;">
            <p><strong>Open Function Server Maintenance</strong></p>
            <nav>
                <ul>
                    <a href="#">Dashboard</a>
                    <a href="sites">Sites</a>
                </ul>
            </nav>
        </div>
        <div class="col-10 main-content">
