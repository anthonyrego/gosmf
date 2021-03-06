package shader

func LoadBasicShaders() {

	New(
		"texture",
		`
    #version 330

    uniform mat4 projection;
    uniform mat4 camera;
    uniform mat4 model;
    uniform vec4 color = vec4(1,1,1,1);

    in vec3 vert;
    in vec2 vertTexCoord;
    out vec2 fragTexCoord;
    out vec4 colorV;

    void main() {
        fragTexCoord = vertTexCoord;
        colorV = color;
        gl_Position = projection * camera *  model * vec4(vert, 1);
    }
    `+"\x00",
		`
    #version 330
    uniform sampler2D tex; //this is the texture
    in vec2 fragTexCoord; //this is the texture coord
    in vec4 colorV;
    out vec4 finalColor; //this is the output color of the pixel

    void main() {
        finalColor = texture(tex, fragTexCoord) * colorV;
    }
    `+"\x00",
		[]string{
			"projection",
			"camera",
			"model",
			"color",
		},
	)

	New(
		"color",
		`
    #version 330

    uniform mat4 projection;
    uniform mat4 camera;
    uniform mat4 model;
    uniform vec4 color;

    in vec3 vert;
    out vec4 colorV;

    void main() {
        colorV = color;
        gl_Position = projection * camera *  model * vec4(vert, 1);
    }
    `+"\x00",
		`
    #version 330
    in vec4 colorV;
    out vec4 finalColor;

    void main() {
        finalColor = colorV;
    }
    `+"\x00",
		[]string{
			"projection",
			"camera",
			"model",
			"color",
		},
	)
}
