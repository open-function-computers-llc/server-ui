=====
Pagetitle: Site Details
BodyClasses: layout-details
=====
<h1 class="pt-3">{{ .Domain }}</h1>

<nav class="" x-data="{ locked: {{ .Locked }}, domain: '{{ .Domain }}' }">
    <a class="btn btn-primary" href="https://{{ .Domain }}" target="_blank" rel="nofollow noopener">Visit site</a>
    <a class="btn btn-primary" href="lock/{{ .Domain }}" rel="nofollow noopener" x-show="!locked">Lock Site</a>
    <a class="btn btn-primary" href="unlock/{{ .Domain }}" rel="nofollow noopener" x-show="locked">Unlock Site</a>
    <a class="btn btn-primary" href="#" target="_blank" rel="nofollow noopener">Duplicate Site</a>
</nav>

<p class="mt-3 mb-0">Traffic:</p>
<ul class="nav nav-tabs" id="siteTabs" role="tablist">
    <li class="nav-item" role="presentation">
        <a href="#yesterday" data-bs-toggle="tab" role="tab" aria-controls="yesterday" aria-selected="true" class="nav-link active">Yesterday</a>
    </li>
    <li class="nav-item">
        <a href="#thirty-days" data-bs-toggle="tab" role="tab" aria-controls="thirty-days" aria-selected="false" class="nav-link">Last 30 Days</a>
    </li>
    <li class="nav-item">
        <a href="#full-report" data-bs-toggle="tab" role="tab" aria-controls="full-report" aria-selected="false" class="nav-link">Full Log</a>
    </li>
</ul>
<div class="tab-content" id="tabContent">
    <div class="tab-pane fade show active" id="yesterday" role="tabpanel" aria-labelledby="yesterday">
        <iframe style="width: 100%; height: 80vh;" src="{{ .RoutePrefix }}/stats/{{ .Domain }}/1"></iframe>
    </div>
    <div class="tab-pane fade" id="thirty-days" role="tabpanel" aria-labelledby="thirty-days">
        <iframe style="width: 100%; height: 80vh;" src="{{ .RoutePrefix }}/stats/{{ .Domain }}/30"></iframe>
    </div>
    <div class="tab-pane fade" id="full-report" role="tabpanel" aria-labelledby="full-report">
        <iframe style="width: 100%; height: 80vh;" src="{{ .RoutePrefix }}/stats/{{ .Domain }}/90"></iframe>
    </div>
</div>
