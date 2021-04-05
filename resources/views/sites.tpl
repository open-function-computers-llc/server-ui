=====
Pagetitle: Websites
BodyClasses: layout-dashboard
=====
<h1 class="pt-3">Websites on this server:</h1>

<p>Sites</p>
<ul>
    {{ range $s := .Sites }}
        <li><a href="details/{{ $s }}">{{ $s }}</a></li>
    {{ end }}
</ul>
