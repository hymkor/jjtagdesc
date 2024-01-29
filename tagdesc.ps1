Set-PSDebug -Strict
$count = 0
$result = $null
    $count++
    if ( $count % 2 -ne 0 ){
jj log --no-graph -r 'latest(tags()):: ~ description(exact:"")' | ForEach-Object {
        $fields = ($_ -split " +")
        if ( $count -eq 1 ){
            $current = $fields[0]
        }
        if ( $fields.Length -eq 7 ){
            if ( $count -eq 1 ){
                $result = $fields[5]
            } else {
                $result = ("{0}-{1}-{2}" -f $fields[5],$count,$current)
            }
        }
        $ncommits++
    }
}
if ( -not $result ){
    exit 1
}
Write-Output $result
