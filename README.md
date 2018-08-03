# tensor-flow-inception-client
Simple grpc client for tensorflow with inception in Go.

If you have just up tensorflow docker locally and want to do inference against inception model then 

Run command from root mentioned below 

`
~/src/github.com/itnilesh/tensor-flow-inception-client/bin/secure

`
`
bash-3.2$ ./inception_secure_client  --serving-address localhost:9000  --image-file-path ./../../clients/resources/models/inception/image_54_mb.jpg
 passed flags are >>  --serving-address localhost:9000  --auth-token  --image-file-path ./../../clients/resources/models/inception/image_54_mb.jpg --cert-file-path  2018/08/03 18:19:26 Warn: --auth-token flag  is not passed
2018/08/03 18:19:26 Warn: --cert-file-path flag is not passed
2018/08/03 18:19:26 Calling predict..
2018/08/03 18:19:28 Predict took 2.404787679s
2018/08/03 18:19:28 Got response
2018/08/03 18:19:28 model_spec:<name:"inception" version:<value:1 > signature_name:"predict_images" > outputs:<key:"classes" value:<dtype:DT_STRING tensor_shape:<dim:<size:1 > dim:<size:5 >> string_val:"great white shark, white shark, man-eater, man-eating shark, Carcharodon carcharias" string_val:"grey whale, gray whale, devilfish, Eschrichtius gibbosus, Eschrichtius robustus" string_val:"breakwater, groin, groyne, mole, bulwark, seawall, jetty" string_val:"killer whale, killer, orca, grampus, sea wolf, Orcinus orca" string_val:"promontory, headland, head, foreland" > > outputs:<key:"scores" value:<dtype:DT_FLOAT tensor_shape:<dim:<size:1 > dim:<size:5 > > float_val:6.625527 float_val:6.1647816 float_val:5.84122 float_val:5.4420643 float_val:5.358542 > >
`
