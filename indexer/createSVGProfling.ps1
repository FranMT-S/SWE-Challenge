Set-Location $PSScriptRoot

gci "profiling" | ? {$_.Extension -like "*prof"} |  % {

    
    if(-not (Test-Path "profiling/$($_.Basename).svg")){
       go tool pprof -svg "profiling/$($_.Name)"  > "profiling/$($_.Basename).svg"
    }
    
   

}