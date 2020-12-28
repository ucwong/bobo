for i in {1..1000}
do
	curl -X POST -d "some data "$i http://localhost:8080/v2/example?k=1\&v=$i
	echo ""
	curl http://localhost:8080/v2/example?k=1
	echo ""
done
