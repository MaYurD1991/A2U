# A2U

FOR UNICODE TO ASCII
go run mail.go -i=input -o=output -c=a 

FOR ASCII TO UNICODE
go run mail.go -i=input -o=output -c=u

BY DEFAULT WORKERS:5

IF YOUR SYSTEM SUPPORT MORE WORKER THEN

FOR UNICODE TO ASCII
go run mail.go -i=input -o=output -c=a -w=10

FOR ASCII TO UNICODE
go run mail.go -i=input -o=output -c=u -w=10

-i = <input file (if not in same directory then including path)>

-o = <output file (if not in same directory then including path)>

-c = a for ASCII / u for UNICODE

-w = <no of workers>

FOR HELP
go run mail.go -h
