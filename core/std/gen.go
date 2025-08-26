//go:generate go run ./internal/generators/letter -pkg std -name RGB -nameL rgb -cmpts R,G,B -cmptsL r,g,b -out rgb.go
//go:generate go run ./internal/generators/letter -pkg std -name RGBA -nameL rgba -cmpts R,G,B,A -cmptsL r,g,b,a -out rgba.go

//go:generate go run ./internal/generators/letter -pkg std -name HSB -nameL hsb -cmpts H,S,B -cmptsL h,s,b -out hsb.go
//go:generate go run ./internal/generators/letter -pkg std -name HSV -nameL hsv -cmpts H,S,V -cmptsL h,s,v -out hsv.go
//go:generate go run ./internal/generators/letter -pkg std -name HSL -nameL hsl -cmpts H,S,L -cmptsL h,s,l -out hsl.go

//go:generate go run ./internal/generators/letter -pkg std -name CMYK -nameL cmyk -cmpts C,M,Y,K -cmptsL c,m,y,k -out cmyk.go
//go:generate go run ./internal/generators/letter -pkg std -name YCbCr -nameL yCbCr -cmpts Y,Cb,Cr -cmptsL y,cb,cr -out yCbCr.go

//go:generate go run ./internal/generators/letter -pkg std -name ST -nameL st -cmpts S,T -cmptsL s,t -out st.go
//go:generate go run ./internal/generators/letter -pkg std -name UV -nameL uv -cmpts U,V -cmptsL u,v -out uv.go
//go:generate go run ./internal/generators/letter -pkg std -name UVW -nameL uvw -cmpts U,V,W -cmptsL u,v,w -out uvw.go

//go:generate go run ./internal/generators/letter -pkg std -name XY -nameL xy -cmpts X,Y -cmptsL x,y -out xy.go
//go:generate go run ./internal/generators/letter -pkg std -name XYZ -nameL xyz -cmpts X,Y,Z -cmptsL x,y,z -out xyz.go
//go:generate go run ./internal/generators/letter -pkg std -name XYZW -nameL xyzw -cmpts X,Y,Z,W -cmptsL x,y,z,w -out xyzw.go
//go:generate go run ./internal/generators/letter -pkg std -name XYZWA -nameL xyzwa -cmpts X,Y,Z,W,A -cmptsL x,y,z,w,a -out xyzwa.go
//go:generate go run ./internal/generators/letter -pkg std -name XYZWAB -nameL xyzwab -cmpts X,Y,Z,W,A,B -cmptsL x,y,z,w,a,b -out xyzwab.go
//go:generate go run ./internal/generators/letter -pkg std -name XYZWABC -nameL xyzwabc -cmpts X,Y,Z,W,A,B,C -cmptsL x,y,z,w,a,b,c -out xyzwabc.go
package std
