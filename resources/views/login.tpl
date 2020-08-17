=====
Pagetitle: Login
BodyClasses: layout-login
=====
<h1>Please Login</h1>
<form action="{{ .RoutePrefix }}/handle-login" method="POST">
    <input type="text" name="username" placeholder="Username" autofocus />
    <input type="password" name="password" placeholder="Password" />
    <input type="submit" value="Login" />
</form>
