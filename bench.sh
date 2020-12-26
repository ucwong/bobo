for i in {1..1000}
do
	curl http://localhost:8080/set?k=1\&v=$i
	echo ""
	curl http://localhost:8080/get?k=1
	echo ""
done
