for i in {1..1}
do
	curl -X POST -d "example info" http://localhost:8080/v2/example/0x2a2A0667f9cBF4055E48eaF0d5B40304B8822184?msg=aHellox\&sig=0xee78eaa27526b412d0e970b85f47c96aa0aa67ed1c06f577ffe712a91284659a0a38529194a53891c84919369e09bf7e08d1655544cb044671461e210ddad1eb00
	echo ""
	#curl http://localhost:8080/v2/example/0xa19d069d48d2e9392ec2bb41ecab0a72119d633b
	#echo ""
done
