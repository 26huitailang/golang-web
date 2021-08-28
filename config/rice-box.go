package config

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file3 := &embedded.EmbeddedFile{
		Filename:    "error/400.html",
		FileModTime: time.Unix(1584145795, 0),

		Content: string("<!DOCTYPE html>\n<html lang=\"en\">\n\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n    <!-- Bootstrap CSS -->\n    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\"\n        integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\" crossorigin=\"anonymous\">\n    <title>Hello, world!</title>\n</head>\n\n<body>\n    <div class=\"container\">\n        <div class=\"row\">\n            <h1 class=\"mx-auto\">Oops! 401</h1>\n        </div>\n        <div class=\"row\">\n            <a class=\"btn btn-primary  mx-auto\" href=\"/\" role=\"button\">HomePage</a>\n        </div>\n    </div>\n    <!-- Optional JavaScript -->\n    <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n    <script src=\"https://code.jquery.com/jquery-3.3.1.slim.min.js\"\n        integrity=\"sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo\"\n        crossorigin=\"anonymous\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js\"\n        integrity=\"sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1\"\n        crossorigin=\"anonymous\"></script>\n    <script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js\"\n        integrity=\"sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM\"\n        crossorigin=\"anonymous\"></script>\n</body>\n\n</html>\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "error/401.html",
		FileModTime: time.Unix(1584145795, 0),

		Content: string("<!DOCTYPE html>\n<html lang=\"en\">\n\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n    <!-- Bootstrap CSS -->\n    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\"\n        integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\" crossorigin=\"anonymous\">\n    <title>Hello, world!</title>\n</head>\n\n<body>\n    <div class=\"container\">\n        <div class=\"row\">\n            <h1 class=\"mx-auto\">Oops! 400</h1>\n        </div>\n        <div class=\"row\">\n            <a class=\"btn btn-primary  mx-auto\" href=\"/\" role=\"button\">HomePage</a>\n        </div>\n    </div>\n    <!-- Optional JavaScript -->\n    <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n    <script src=\"https://code.jquery.com/jquery-3.3.1.slim.min.js\"\n        integrity=\"sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo\"\n        crossorigin=\"anonymous\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js\"\n        integrity=\"sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1\"\n        crossorigin=\"anonymous\"></script>\n    <script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js\"\n        integrity=\"sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM\"\n        crossorigin=\"anonymous\"></script>\n</body>\n\n</html>\n"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "error/404.html",
		FileModTime: time.Unix(1584145795, 0),

		Content: string("<!DOCTYPE html>\n<html lang=\"en\">\n\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n    <!-- Bootstrap CSS -->\n    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\"\n        integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\" crossorigin=\"anonymous\">\n    <title>Hello, world!</title>\n</head>\n\n<body>\n    <div class=\"container\">\n        <div class=\"row\">\n            <h1 class=\"mx-auto\">Oops! 404 Not Found</h1>\n        </div>\n        <div class=\"row\">\n            <a class=\"btn btn-primary  mx-auto\" href=\"/\" role=\"button\">HomePage</a>\n        </div>\n    </div>\n    <!-- Optional JavaScript -->\n    <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n    <script src=\"https://code.jquery.com/jquery-3.3.1.slim.min.js\"\n        integrity=\"sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo\"\n        crossorigin=\"anonymous\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js\"\n        integrity=\"sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1\"\n        crossorigin=\"anonymous\"></script>\n    <script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js\"\n        integrity=\"sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM\"\n        crossorigin=\"anonymous\"></script>\n</body>\n\n</html>\n"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "layouts/layout.html",
		FileModTime: time.Unix(1594047687, 0),

		Content: string("{{define \"layout\"}}\n<!DOCTYPE html>\n<html lang=\"en\">\n\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n    <link\n        href=\"data:image/x-icon;base64,AAABAAEAEBAAAAEACABoBQAAFgAAACgAAAAQAAAAIAAAAAEACAAAAAAAAAEAAAAAAAAAAAAAAAEAAAAAAACxSagAAAAAALJJqAC1UK0AkDqBAPn5+QCiQZYAizZ8AK5IowCMNnwAsn6oALRLqwCVRIYAkkuDAJ1SkAC5YLIAtlmtAJM8hACDMXIAmj+MAKxGoQC2Ta4AsE2mALFNpgCWTYYAtVCuAIczdQCtS6QAkj2CAI05fQCqSp8AiTh4AKxGogCmRpoAvWK2AK1JogC3UK8AqEWdAIYzdgCHM3YAkjqDAI82fgC1S60AgjJxALVdrACkVpcAjTl+ALxgtACmQ5sAjjx+AIQ0dACcP44A/v7+AJ1QkACIM3cAqEieAI06fACzT6sAkj2EALRPqwCobJwArkumAIMycgCjXZUAizV6AKFDlACzSqkAkTiCALZRrgCVP4cAmDuKALNNqQCGN3UAhzd1AKpFnwCqV54AkzqFAJQ6hQCwSKcAsFqmALRPrAC1T6wAgzJzAJs9jQCUPYUAhDJzAKFDlQCQPIAAt1GvAKhGnQC5X7EAo1SXAK9JpQCYPosAhzd2AI83fgCIN3YAtkytAIk2eQCLOXkAsU6oALdRsAC4UbAAqUaeAJM7hACwSaYAiDd3AJo+jAC2TK4AlEyGAJlBjACJNnoAvWG1AIYydQCLOXoA////ALJOqQC1SqwAuFGxAI04fQCfP5IAo0aXAI07fQD9/f0AokGVAJ1AkACjZJYAs06qAJhOigD7+/sAnEOQAKRDmACDMXEAm0qNAI07fgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQEBAQEBAQEBAQEBAQEBAQEBbYA1W0tPLFpwIi8PAQEBK1JxQENGeDAgAipsYUIBATI+hBopTVN8JRQAdRVhAQFvJhKCczR7e3MfCE4LFQEBBAcngnM/DIWCNB9KaQsBAV0oCYJzGFU2X3MfBmdcAQFBM0yCcwp+PII0a2t8WQEBN4N9gnOBBQVzH3doE1YBAT0jIYJzDh1ignMxLlRuAQF/RxuCcy0Reg1zHx+GOgEBGSQ5gnNzNDRzHx9JHzgBAVB2ZRAQEBAQEBxjXklqAQFkJHZmURceeYJFV3JgSAEBAXQDWEQ7Fh55gkVXcgEBAQEBAQEBAQEBAQEBAQEBAf//AADAAwAAgAEAAIABAACAAQAAgAEAAIABAACAAQAAgAEAAIABAACAAQAAgAEAAIABAACAAQAAwAMAAP//AAA=\"\n        rel=\"icon\" type=\"image/x-icon\" />\n    <!-- Bootstrap CSS -->\n    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\"\n        integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\" crossorigin=\"anonymous\">\n    <title>Hey!</title>\n</head>\n\n<body style=\"padding-top: 60px\">\n    <nav class=\"navbar fixed-top navbar-expand-lg navbar-light bg-light\">\n        <a class=\"navbar-brand\" href=\"/\">Home</a>\n        <button class=\"navbar-toggler\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarNav\"\n            aria-controls=\"navbarNav\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n            <span class=\"navbar-toggler-icon\"></span>\n        </button>\n        <div class=\"collapse navbar-collapse\" id=\"navbarNav\">\n            <ul class=\"navbar-nav\">\n                <li class=\"nav-item\">\n                    <a class=\"nav-link\" href=\"/suites?is_like=true\">Like</a>\n                </li>\n            </ul>\n            <ul class=\"navbar-nav\">\n                <li class=\"nav-item\">\n                    <a class=\"nav-link\" href=\"/devops\">Devops</a>\n                </li>\n            </ul>\n        </div>\n    </nav>\n    <div class=\"container\">\n        {{template \"content\" .}}\n    </div>\n\n    <!-- Optional JavaScript -->\n    <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n    {{template \"script\"}}\n    <script src=\"https://cdn.bootcdn.net/ajax/libs/jquery/3.3.1/jquery.slim.min.js\"\n        integrity=\"sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo\"\n        crossorigin=\"anonymous\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js\"\n        integrity=\"sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1\"\n        crossorigin=\"anonymous\"></script>\n    <script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js\"\n        integrity=\"sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM\"\n        crossorigin=\"anonymous\"></script>\n</body>\n\n</html>\n{{end}}"),
	}
	file9 := &embedded.EmbeddedFile{
		Filename:    "pages/devops.html",
		FileModTime: time.Unix(1584145795, 0),

		Content: string("{{define \"content\"}}\n<div class=\"row\">\n    <div class=\"col-md-2\"></div>\n    <div class=\"col-xs-12 col-md-8\">\n        <div class=\"row\">\n            <h5 style=\"margin: 0 auto\">Devops</h5>\n        </div>\n        <div class=\"row\">\n            <form action=\"/devops/initdb\" method=\"POST\">\n                <div class=\"form-group\">\n                    <label for=\"csrf\">重置数据库</label>\n                    <!-- <input type=\"password\" class=\"form-control\" id=\"exampleInputPassword1\" placeholder=\"Password\"> -->\n                    <input type=\"hidden\" name=\"csrf\" value=\"{{.}}\">\n                </div>\n                <!-- <input type=\"submit\"> -->\n                <button type=\"submit\" class=\"btn btn-primary\">Submit</button>\n            </form>\n        </div>\n        <div class=\"row\">\n            <form action=\"/task/theme\" method=\"POST\">\n                <div class=\"form-group\">\n                    <label for=\"url\">Theme</label>\n                    <input type=\"text\" class=\"form-control\" id=\"url\" name=\"url\"\n                        placeholder=\"https://www.meituri.com/x/85/\">\n                    <input type=\"hidden\" name=\"csrf\" value=\"{{.}}\">\n                </div>\n                <button type=\"submit\" class=\"btn btn-primary\">Submit</button>\n            </form>\n        </div>\n        <div class=\"row\">\n            <form action=\"/task/suite\" method=\"POST\">\n                <div class=\"form-group\">\n                    <label for=\"url\">Suite</label>\n                    <input type=\"text\" class=\"form-control\" id=\"url\" name=\"url\"\n                        placeholder=\"https://www.meituri.com/a/27606/\">\n                    <input type=\"hidden\" name=\"csrf\" value=\"{{.}}\">\n                </div>\n                <button type=\"submit\" class=\"btn btn-primary\">Submit</button>\n            </form>\n        </div>\n    </div>\n    <div class=\"col-md-2\"></div>\n</div>\n{{end}}\n{{define \"script\"}}{{end}}\n"),
	}
	filea := &embedded.EmbeddedFile{
		Filename:    "pages/suite.html",
		FileModTime: time.Unix(1584145795, 0),

		Content: string("{{define \"content\"}}\n<div class=\"row\">\n    <div class=\"col-md-2\"></div>\n    <div class=\"col-xs-12 col-md-8\">\n        <div class=\"row\">\n            <h5 style=\"margin: 0 auto\">{{.Suite.Name}}</h5>\n        </div>\n        <div class=\"row\">\n            <ul class=\"list-group\">\n                {{range .Images}}\n                <li class=\"list-group-item\"><img src=\"/image/{{.Path}}\" style=\"width: 100%\" alt=\"?\" /></li>\n                {{end}}\n            </ul>\n        </div>\n        <div class=\"row\">\n            <nav class=\"navbar fixed-bottom navbar-light bg-light\" style=\"display: inline-block\">\n                <a href=\"/suites/{{.Suite.ID}}/doread\" type=\"button\" class=\"btn btn-outline-primary btn-sm\"\n                    name=\"action-read\">\n                    {{if .Suite.IsRead}}\n                    已读\n                    {{else}}\n                    DoRead\n                    {{end}}\n                </a>\n                <a href=\"/suites/{{.Suite.ID}}/dolike\" type=\"button\" class=\"btn btn-outline-danger btn-sm\"\n                    name=\"action-like\">\n                    {{if .Suite.IsLike}}\n                    😍\n                    {{else}}\n                    DoLike\n                    {{end}}\n                </a>\n            </nav>\n        </div>\n    </div>\n    <div class=\"col-md-2\"></div>\n</div>\n{{end}}\n{{define \"script\"}}{{end}}\n"),
	}
	fileb := &embedded.EmbeddedFile{
		Filename:    "pages/suites.html",
		FileModTime: time.Unix(1584145795, 0),

		Content: string("{{define \"content\"}}\n<div class=\"row\">\n    <div class=\"col-md-2\"></div>\n    <div class=\"col-xs-12 col-md-8\">\n        <table class=\"table\">\n            <thead>\n                <tr>\n                    <th scope=\"col\">#</th>\n                    <th scope=\"col\">Read</th>\n                    <th scope=\"col\">Like</th>\n                    <th scope=\"col\">Name</th>\n                </tr>\n            </thead>\n            <tbody>\n                {{range $index, $suite := .}}\n                <tr>\n                    <th scope=\"row\">{{$index}}</th>\n                    <td>{{if $suite.IsRead}}已读{{else}}{{end}}</td>\n                    <td>{{if $suite.IsLike}}😍{{else}}{{end}}</td>\n                    <td><a href=\"/suites/{{$suite.ID}}\">{{$suite.Name}}</a></td>\n                </tr>\n                {{end}}\n            </tbody>\n        </table>\n    </div>\n    <div class=\"col-md-2\"></div>\n</div>\n{{end}}\n{{define \"script\"}}{{end}}\n"),
	}
	filec := &embedded.EmbeddedFile{
		Filename:    "pages/theme.html",
		FileModTime: time.Unix(1584145795, 0),

		Content: string("{{define \"content\"}}\n<div class=\"row\">\n    <div class=\"col-md-2\"></div>\n    <div class=\"col-xs-12 col-md-4\">\n        <p class=\"h2\"><a href=\"/themes/{{.Theme.ID}}\">{{.Theme.Name}}</a></p>\n    </div>\n    <div class=\"col-xs-12 col-md-4\">\n        <a href=\"{{.Theme.ID}}?is_read=false\" class=\"badge badge-primary\">\n            未读 <span class=\"badge badge-light\">{{.CountUnread}}</span>\n        </a>\n        <a href=\"{{.Theme.ID}}?is_read=true\" class=\"badge badge-secondary\">\n            已读 <span class=\"badge badge-light\">{{.CountRead}}</span>\n        </a>\n    </div>\n    <div class=\"col-md-2\"></div>\n</div>\n<div class=\"row\">\n    <div class=\"col-md-2\"></div>\n    <div class=\"col-xs-12 col-md-8\">\n        <table class=\"table\">\n            <thead>\n                <tr>\n                    <th scope=\"col\">#</th>\n                    <th scope=\"col\">Read</th>\n                    <th scope=\"col\">Like</th>\n                    <th scope=\"col\">Name</th>\n                </tr>\n            </thead>\n            <tbody>\n                {{range $index, $suite := .Suites}}\n                <tr>\n                    <th scope=\"row\">{{$index}}</th>\n                    <td>{{if $suite.IsRead}}已读{{else}}{{end}}</td>\n                    <td>{{if $suite.IsLike}}😍{{else}}{{end}}</td>\n                    <td><a href=\"/suites/{{$suite.ID}}\">{{$suite.Name}}</a></td>\n                </tr>\n                {{end}}\n            </tbody>\n        </table>\n    </div>\n    <div class=\"col-md-2\"></div>\n</div>\n{{end}}\n{{define \"script\"}}{{end}}\n"),
	}
	filed := &embedded.EmbeddedFile{
		Filename:    "pages/themes.html",
		FileModTime: time.Unix(1584145795, 0),

		Content: string("{{define \"content\"}}\n<div class=\"row\">\n    <div class=\"col-md-2\"></div>\n    <div class=\"col-xs-12 col-md-8\">\n        <ul class=\"list-group\">\n            {{range .}}\n            <li class=\"list-group-item\"><a href=\"themes/{{.ID}}\">{{.Name}}</a></li>\n            {{end}}\n        </ul>\n    </div>\n    <div class=\"col-md-2\"></div>\n</div>\n{{end}}\n{{define \"script\"}}{{end}}"),
	}
	filee := &embedded.EmbeddedFile{
		Filename:    "pages/websocket.html",
		FileModTime: time.Unix(1584145795, 0),

		Content: string("{{define \"content\"}}\n<p id=\"output\"></p>\n{{end}}\n{{define \"script\"}}\n<script>\n    var loc = window.location;\n    var uri = 'ws:';\n\n    if (loc.protocol === 'https:') {\n        uri = 'wss:';\n    }\n    uri += '//' + loc.host;\n    uri += loc.pathname + '/ws';\n\n    ws = new WebSocket(uri)\n\n    ws.onopen = function () {\n        console.log('Connected')\n    }\n\n    ws.onmessage = function (evt) {\n        var out = document.getElementById('output');\n        out.innerHTML += evt.data + '<br>';\n    }\n\n    setInterval(function () {\n        ws.send('Hello, Server!');\n    }, 1000);\n</script>\n{{end}}\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1584145795, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dir2 := &embedded.EmbeddedDir{
		Filename:   "error",
		DirModTime: time.Unix(1584145795, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file3, // "error/400.html"
			file4, // "error/401.html"
			file5, // "error/404.html"

		},
	}
	dir6 := &embedded.EmbeddedDir{
		Filename:   "layouts",
		DirModTime: time.Unix(1594047687, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file7, // "layouts/layout.html"

		},
	}
	dir8 := &embedded.EmbeddedDir{
		Filename:   "pages",
		DirModTime: time.Unix(1584145795, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file9, // "pages/devops.html"
			filea, // "pages/suite.html"
			fileb, // "pages/suites.html"
			filec, // "pages/theme.html"
			filed, // "pages/themes.html"
			filee, // "pages/websocket.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir2, // "error"
		dir6, // "layouts"
		dir8, // "pages"

	}
	dir2.ChildDirs = []*embedded.EmbeddedDir{}
	dir6.ChildDirs = []*embedded.EmbeddedDir{}
	dir8.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`../templates`, &embedded.EmbeddedBox{
		Name: `../templates`,
		Time: time.Unix(1584145795, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":        dir1,
			"error":   dir2,
			"layouts": dir6,
			"pages":   dir8,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"error/400.html":       file3,
			"error/401.html":       file4,
			"error/404.html":       file5,
			"layouts/layout.html":  file7,
			"pages/devops.html":    file9,
			"pages/suite.html":     filea,
			"pages/suites.html":    fileb,
			"pages/theme.html":     filec,
			"pages/themes.html":    filed,
			"pages/websocket.html": filee,
		},
	})
}
