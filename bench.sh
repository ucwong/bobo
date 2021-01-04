#!/bin/sh
domain="http://share.cortexlabs.ai:8081"
for i in {1..1}
do
	echo "Register a user detail ."
	curl -X POST -d "{\"ts\":1609310997, \"name\":\"jo\"}" http://localhost:8080/user/0x970E8128AB834E8EAC17Ab8E3812F010678CF791?sig=0x15ce17f60e6825a4d5556867c30d3bc823f9f2dd0d55aa845a816f4518a081ca5e2c9fea9ec552e861d015306c6c7c4132135e97b0e695e01c751c51e5e7075d01
	echo ""
	#echo "Update a user detail ."
	curl -X POST -d "{\"ts\":1609310997, \"name\":\"jo\", \"age\":30}" ${domain}/user/0x970E8128AB834E8EAC17Ab8E3812F010678CF791?sig=0x2e9d7610f3611be41ded12d595285530387522fbc60ca691811adb4306996bc6275e80e292041dd4eecc65bc1d6e53e410eb7390e0fca6fd7d2c378a6f3efa7000
	echo ""
	#echo "Favor picture 0x2a2a0667f9cbf4055e48eaf0d5b40304b8822184 ."
	curl -X POST -d "{\"ts\":1609310997, \"addr\":\"0x2a2a0667f9cbf4055e48eaf0d5b40304b8822184\"}" http://localhost:8080/favor/0x970E8128AB834E8EAC17Ab8E3812F010678CF791?sig=0xcec3911c55e99dff3b6174fa91e8856f1374c58ca239801201797806f4a0c6355b4fc7725ca39b54ff5eb50b86f013c0e5c06c28d8b8fe0a9fbe25e27e26bfa601
	echo ""
	#echo "Favor picture 0x970E8128AB834E8EAC17Ab8E3812F010678CF791 ."
	curl -X POST -d "{\"ts\":1609310997, \"addr\":\"0x970E8128AB834E8EAC17Ab8E3812F010678CF791\"}" http://localhost:8080/favor/0x970E8128AB834E8EAC17Ab8E3812F010678CF791?sig=0xdf9e25d0da49305c53dff42519a9c9c3a02c4f29a2c15333c7b403ec9ae5dcb10bf12598441c7bc53ba4dc66a85bc77440ce61c72e2ab76a43f38a34345ce4ce00
	echo ""
	#echo "Favor picture 0x564286362092d8e7936f0549571a803b203aaced ."
	curl -X POST -d "{\"ts\":1609310997, \"addr\":\"0x564286362092d8e7936f0549571a803b203aaced\"}" http://localhost:8080/favor/0x970E8128AB834E8EAC17Ab8E3812F010678CF791?sig=0x6cb41f30e9dc732c4fc01cd3288c020a7473be57510a63dcd9d696836f2495ff7c0544fca083776f17d050b93d92e140323eb1c3c179e7e3b830db710afddc3e01
	echo ""
	#echo "Get 0x970e8128ab834e8eac17ab8e3812f010678cf791 user detail ."
	curl -X GET http://share.cortexlabs.ai:8081/user/0x970e8128ab834e8eac17ab8e3812f010678cf791
	echo ""
	#echo "Get 0x970e8128ab834e8eac17ab8e3812f010678cf791 favor pictures ."
	curl -X GET http://share.cortexlabs.ai:8081/favor/0x970e8128ab834e8eac17ab8e3812f010678cf791
	echo ""
	#echo "Get picture 0x970e8128ab834e8eac17ab8e3812f010678cf791 fans ."
	curl -X GET http://share.cortexlabs.ai:8081/favored/0x970e8128ab834e8eac17ab8e3812f010678cf791
	echo ""
	#echo "Delete a favor picture 0x564286362092d8e7936f0549571a803b203aaced ."
	curl -X DELETE -d "{\"ts\":1609310997, \"addr\":\"0x564286362092d8e7936f0549571a803b203aaced\"}" http://localhost:8080/favor/0x970E8128AB834E8EAC17Ab8E3812F010678CF791?sig=0x6cb41f30e9dc732c4fc01cd3288c020a7473be57510a63dcd9d696836f2495ff7c0544fca083776f17d050b93d92e140323eb1c3c179e7e3b830db710afddc3e01
	echo ""
	#echo "Get 0x970e8128ab834e8eac17ab8e3812f010678cf791 favor pictures again ."
	curl -X GET http://share.cortexlabs.ai:8081/favor/0x970e8128ab834e8eac17ab8e3812f010678cf791
	echo ""
	curl -X POST -d "{\"ts\":1609310997, \"addr\":\"0x564286362092d8e7936f0549571a803b203aaced\"}" http://localhost:8080/favor/0x970E8128AB834E8EAC17Ab8E3812F010678CF791?sig=0x6cb41f30e9dc732c4fc01cd3288c020a7473be57510a63dcd9d696836f2495ff7c0544fca083776f17d050b93d92e140323eb1c3c179e7e3b830db710afddc3e01
	echo ""
	curl -X GET http://share.cortexlabs.ai:8081/favor/0x970e8128ab834e8eac17ab8e3812f010678cf791
	echo ""
	#echo "follow some one"
	curl -X POST -d "{\"ts\":1609310997, \"addr\":\"0x2a2a0667f9cbf4055e48eaf0d5b40304b8822184\"}" http://localhost:8080/follow/0x970E8128AB834E8EAC17Ab8E3812F010678CF791?sig=0xcec3911c55e99dff3b6174fa91e8856f1374c58ca239801201797806f4a0c6355b4fc7725ca39b54ff5eb50b86f013c0e5c06c28d8b8fe0a9fbe25e27e26bfa601
	echo ""
	curl -X POST -d "{\"ts\":1609310997, \"addr\":\"0x564286362092d8e7936f0549571a803b203aaced\"}" http://localhost:8080/follow/0x970E8128AB834E8EAC17Ab8E3812F010678CF791?sig=0x6cb41f30e9dc732c4fc01cd3288c020a7473be57510a63dcd9d696836f2495ff7c0544fca083776f17d050b93d92e140323eb1c3c179e7e3b830db710afddc3e01
	echo ""
	#echo "follow list"
	curl -X GET http://share.cortexlabs.ai:8081/follow/0x970e8128ab834e8eac17ab8e3812f010678cf791
	echo ""
	#echo "fans"
	curl -X GET http://share.cortexlabs.ai:8081/followed/0x564286362092d8e7936f0549571a803b203aaced
	echo ""
	curl -X GET http://share.cortexlabs.ai:8081/followed/0x2a2a0667f9cbf4055e48eaf0d5b40304b8822184
	echo ""
	#echo "unfollow"
	curl -X DELETE -d "{\"ts\":1609310997, \"addr\":\"0x564286362092d8e7936f0549571a803b203aaced\"}" http://localhost:8080/follow/0x970E8128AB834E8EAC17Ab8E3812F010678CF791?sig=0x6cb41f30e9dc732c4fc01cd3288c020a7473be57510a63dcd9d696836f2495ff7c0544fca083776f17d050b93d92e140323eb1c3c179e7e3b830db710afddc3e01
	echo ""
	curl -X GET http://share.cortexlabs.ai:8081/followed/0x564286362092d8e7936f0549571a803b203aaced
	echo ""
	echo "Finish"
done
