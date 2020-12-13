<!doctype html>
<html class="no-js" lang="">

<head>
    <meta charset="utf-8">
    <title>badpod</title>
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.8.2/css/bulma.min.css">
    <!--<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma-extensions@6.2.7/bulma-switch/dist/css/bulma-switch.min.css" integrity="sha256-hhNzSX9QCUNRpgKiGuOGzPtUdetKhSP4X/jQkkYgBzI=" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bulma-extensions@6.2.7/dist/js/bulma-extensions.min.js" integrity="sha256-q4zsxO0fpPm6VhtL/9QkCFE5ZkNa0yeUxhmt1VO1ev0=" crossorigin="anonymous"></script>-->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css">
    <meta name="theme-color" content="#fafafa">
</head>

<body>

<section class="section">
    <div class="container">
        <h1 class="title is-1">badpod</h1>
        <h2 class="subtitle">{{.Version}}, git {{.GitCommit}}, built {{.BuildTime}}</h2>

        <h1 class="title is-3">Session</h1>
        <p>
            Started at {{.StartTime}}; up {{.RunTime}}<br>
            Name: {{.SessionName}}<br>
            Request: {{.RequestNumber}}<br>
        </p>

        <h1 class="title is-3">Resources</h1>
        <h1 class="title is-5">CPU</h1>
        <p>
            Apparent cores: {{.VirtCores}}<br>
            Usage time: {{.CpuSelfTime}}<br>
        </p>

        <h1 class="title is-5">Memory</h1>
        <p>
            Apparent size: {{.MemTotal}}<br>
            Usage: physical {{.MemUsePhysical}}; virtual {{.MemUseVirtualTotal}} (of which {{.MemUseVirtualRuntime}} go runtime)<br>
            GC Runs: {{.GcRuns}}<br>
        </p>
    </div>
</section>

<section class="section">

    <div class="field is-horizontal">
        <div class="field-label is-normal">
            <label class="label">Quit</label>
        </div>
        <div class="field-body">
            <form method="get" action="handlers/exit">
                <div class="field has-addons">
                    <p class="control has-icons-left">
                        <input class="input" type="text" name="code" value="0" placeholder="Exit Code">
                        <span class="icon is-small is-left"><i class="fas fa-skull"></i></span>
                    </p>
                    <p class="control">
                        <button class="button">Exit</button>
                    </p>
                </div>
            </form>
        </div>
    </div>

    <div class="field is-horizontal">
        <div class="field-label is-normal">
            <label class="label">Liveness</label>
        </div>
        <div class="field-body">
            <div class="field has-addons">
                <p class="control">
                    <form method="get" action="handlers/liveness">
                        <button class="button {{if eq .SettingLiveness "true"}}is-primary{{end}}" name="value" value="true">true</button>
                        <button class="button {{if eq .SettingLiveness "false"}}is-primary{{end}}" name="value" value="false">false</button>
                    </form>
                </p>
            </div>
        </div>
    </div>

    <div class="field is-horizontal">
        <div class="field-label is-normal">
            <label class="label">Readiness</label>
        </div>
        <div class="field-body">
            <div class="field has-addons">
                <div class="control">
                    <form method="get" action="handlers/readiness">
                        <button class="button {{if eq .SettingReadiness "true"}}is-primary{{end}}" name="value" value="true">true</button>
                        <button class="button {{if eq .SettingReadiness "false"}}is-primary{{end}}" name="value" value="false">false</button>
                    </form>
                </div>
            </div>
        </div>
    </div>

    <div class="field is-horizontal">
        <div class="field-label is-normal">
            <label class="label">Latency</label>
        </div>
        <div class="field-body">
            <form method="get" action="handlers/delay">
                <div class="field has-addons">
                    <p class="control has-icons-left">
                        <input class="input" type="text" name="value" value="{{.SettingLatency}}" placeholder="s">
                        <span class="icon is-small is-left"><i class="fas fa-user-clock"></i></span>
                    </p>
                    <p class="control">
                        <button class="button">Set</button>
                    </p>
                </div>
            </form>
        </div>
    </div>

    <div class="field is-horizontal">
        <div class="field-label is-normal">
            <label class="label">Bandwidth</label>
        </div>
        <div class="field-body">
            <form method="get" action="handlers/bandwidth">
                <div class="field has-addons">
                    <p class="control has-icons-left">
                        <input class="input" type="text" name="value" value="{{.SettingBandwidth}}" placeholder="bytes/s">
                        <span class="icon is-small is-left"><i class="fas fa-tachometer-alt"></i></span>
                    </p>
                    <p class="control">
                        <button class="button">Set</button>
                    </p>
                </div>
            </form>
        </div>
    </div>

    <div class="field is-horizontal">
        <div class="field-label is-normal">
            <label class="label">Error Rate</label>
        </div>
        <div class="field-body">
            <form method="get" action="handlers/errorrate">
                <div class="field has-addons">
                    <p class="control has-icons-left">
                        <input class="input" type="text" name="value" value="{{.SettingErrorRate}}" placeholder="rate [0-1]">
                        <span class="icon is-small is-left"><i class="fas fa-bomb"></i></span>
                    </p>
                    <p class="control">
                        <button class="button">Set</button>
                    </p>
                </div>
            </form>
        </div>
    </div>

    <div class="field is-horizontal">
        <div class="field-label is-normal">
            <label class="label">Allocate</label>
        </div>
        <div class="field-body">
            <div class="field has-addons">
                <div class="control">
                    <form method="get" action="handlers/allocate">
                        <button class="button" name="value" value="1024">1kB</button>
                        <button class="button" name="value" value="1048576">1MB</button>
                        <button class="button" name="value" value="1073741824">1GB</button>
                    </form>
                </div>
            </div>
        </div>
    </div>

    <div class="field is-horizontal">
        <div class="field-label is-normal">
            <label class="label">CPU Use</label>
        </div>
        <div class="field-body">
            <form method="get" action="handlers/cpu">
                <div class="field has-addons">
                    <p class="control has-icons-left">
                        <input class="input" type="text" name="value" value="{{.SettingCpuUse}}" placeholder="cores [0-n]">
                        <span class="icon is-small is-left"><i class="fas fa-microchip"></i></span>
                    </p>
                    <p class="control">
                        <button class="button">Set</button>
                    </p>
                </div>
            </form>
        </div>
    </div>

</section>

<script src="https://use.fontawesome.com/releases/v5.3.1/js/all.js"></script>

</body>

</html>
