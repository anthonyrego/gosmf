package render

var defaultShaders = map[string]Shader{
	"default": Shader{
		vert: `
  	#version 330

    uniform mat4 projection;
    uniform mat4 camera;
    uniform mat4 model;

  	in vec3 vert;
  	in vec2 vertTexCoord;
  	out vec2 fragTexCoord;

  	void main() {
  	    fragTexCoord = vertTexCoord;

  	    gl_Position = projection * model * vec4(vert, 1);
  	}
    ` + "\x00",
		frag: `
  	#version 330
  	uniform sampler2D tex; //this is the texture
  	in vec2 fragTexCoord; //this is the texture coord
  	out vec4 finalColor; //this is the output color of the pixel

  	void main() {
  	    finalColor = texture(tex, fragTexCoord);
  	}
    ` + "\x00",
		id:     0,
		loaded: false,
	},
}
