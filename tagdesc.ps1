Set-PSDebug -Strict
$nlines = 0
$ncommits = 0
jj log --no-graph -r 'latest(tags()):: ~ description(exact:"")' | ForEach-Object {
    if ( ($nlines++ % 2) -eq 0 ){
        $fields = ($_ -split " +")
        if ( $ncommits -eq 0 ){
            $current = $fields[0]
        }
        if ( $fields.Length -eq 7 ){
            if ( $ncommits -eq 0 ){
                $result = $fields[5]
            } else {
                $result = ("{0}-{1}-{2}" -f $fields[5],$ncommits,$current)
            }
        }
        $ncommits++
    }
}
Write-Output $result
