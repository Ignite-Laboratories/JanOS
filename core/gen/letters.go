/**
Letter Types
*/

//go:generate go run ./letters/main -pkg std -name _cursor -nameL _cursor -cmpts I -cmptsL i -out ../std/cursor_gen.go

//go:generate go run ./letters/main -pkg std -name FMA -nameL fma -cmpts Force,Mass,Acceleration -cmptsL force,mass,acceleration -out ../std/fma.go
//go:generate go run ./letters/main -pkg std -name RGB -nameL rgb -cmpts R,G,B -cmptsL r,g,b -out ../std/rgb.go
//go:generate go run ./letters/main -pkg std -name RGBA -nameL rgba -cmpts R,G,B,A -cmptsL r,g,b,a -out ../std/rgba.go
//go:generate go run ./letters/main -pkg std -name HSB -nameL hsb -cmpts H,S,B -cmptsL h,s,b -out ../std/hsb.go
//go:generate go run ./letters/main -pkg std -name HSV -nameL hsv -cmpts H,S,V -cmptsL h,s,v -out ../std/hsv.go
//go:generate go run ./letters/main -pkg std -name HSL -nameL hsl -cmpts H,S,L -cmptsL h,s,l -out ../std/hsl.go
//go:generate go run ./letters/main -pkg std -name CMYK -nameL cmyk -cmpts C,M,Y,K -cmptsL c,m,y,k -out ../std/cmyk.go
//go:generate go run ./letters/main -pkg std -name YCbCr -nameL yCbCr -cmpts Y,Cb,Cr -cmptsL y,cb,cr -out ../std/yCbCr.go
//go:generate go run ./letters/main -pkg std -name ST -nameL st -cmpts S,T -cmptsL s,t -out ../std/st.go
//go:generate go run ./letters/main -pkg std -name UV -nameL uv -cmpts U,V -cmptsL u,v -out ../std/uv.go
//go:generate go run ./letters/main -pkg std -name UVW -nameL uvw -cmpts U,V,W -cmptsL u,v,w -out ../std/uvw.go
//go:generate go run ./letters/main -pkg std -name XY -nameL xy -cmpts X,Y -cmptsL x,y -out ../std/xy.go
//go:generate go run ./letters/main -pkg std -name XYZ -nameL xyz -cmpts X,Y,Z -cmptsL x,y,z -out ../std/xyz.go
//go:generate go run ./letters/main -pkg std -name XYZW -nameL xyzw -cmpts X,Y,Z,W -cmptsL x,y,z,w -out ../std/xyzw.go
//go:generate go run ./letters/main -pkg std -name XYZWA -nameL xyzwa -cmpts X,Y,Z,W,A -cmptsL x,y,z,w,a -out ../std/xyzwa.go
//go:generate go run ./letters/main -pkg std -name XYZWAB -nameL xyzwab -cmpts X,Y,Z,W,A,B -cmptsL x,y,z,w,a,b -out ../std/xyzwab.go
//go:generate go run ./letters/main -pkg std -name XYZWABC -nameL xyzwabc -cmpts X,Y,Z,W,A,B,C -cmptsL x,y,z,w,a,b,c -out ../std/xyzwabc.go
package gen
