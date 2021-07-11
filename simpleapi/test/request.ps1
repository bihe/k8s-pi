Write-Host "Query the simpleapi @ http://192.168.1.123"
for() {
    $res = Invoke-WebRequest -URI http://192.168.1.123
    $v = $res.Content | jq '.version'
    $h = $res.Content | jq '.hostname'
    $calls += 1

    Start-Sleep -Milliseconds 250
    Write-Host ("`rrequest# ("+$calls+"): V = "+$v+" // hostname = "+$h) -NoNewline
}