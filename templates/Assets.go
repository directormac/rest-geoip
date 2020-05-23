package templates

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetscc9a4ea24b3abb724da85710e7e7309794800cc4 = "{{template \"base\" .}}\n\n{{define \"ipaddress\"}}{{.IP}}{{end}}\n\n{{define \"main\"}}\n<div class=\"container-fluid w-25 p-3 h-100 justify-content-center align-items-center\">\n  <div class=\"card\">\n      <div class=\"card-header\">\n        {{.IP}}\n      </div>\n      <div class=\"card-body\">\n        <table style=\"border-collapse: collapse; border: none;\">\n          <tr style=\"border: none;\">\n            <td style=\"border: none;\">\n              Country\n            </td>\n            <td style=\"border: none;\">\n              {{.Country.ISOCode}}\n            </td>\n          </tr>\n\n          <tr style=\"border: none;\">\n            <td style=\"border: none;\">\n              EU\n            </td>\n            <td style=\"border: none;\">\n              {{if .Country.IsInEuropeanUnion}}\n                Yes\n              {{else}}\n                No\n              {{end}}\n            </td>\n          </tr>\n\n          <tr style=\"border: none;\">\n            <td style=\"border: none;\">\n              City\n            </td>\n            <td style=\"border: none;\">\n              {{range $key, $value := .City.Names}}\n                {{$key}}:{{$value}}</br>\n              {{end}}\n            </td>\n          </tr>\n\n          <tr style=\"border: none;\">\n            <td style=\"border: none;\">\n              Latitude\n            </td>\n            <td style=\"border: none;\">\n              {{.Location.Latitude}}\n            </td>\n          </tr>\n\n          <tr style=\"border: none;\">\n            <td style=\"border: none;\">\n              Longitude\n            </td>\n            <td style=\"border: none;\">\n              {{.Location.Longitude}}\n            </td>\n          </tr>\n\n          <tr style=\"border: none;\">\n            <td style=\"border: none;\">\n              Time Zone\n            </td>\n            <td style=\"border: none;\">\n              {{.Location.TimeZone}}\n            </td>\n          </tr>\n\n          <tr style=\"border: none;\">\n            <td style=\"border: none;\">\n              Approx. Zip Code\n            </td>\n            <td style=\"border: none;\">\n              {{.Postal.Code}}\n            </td>\n          </tr>\n\n          <tr style=\"border: none;\">\n            <td style=\"border: none;\">\n              Anonymous Proxy\n            </td>\n            <td style=\"border: none;\">\n              {{.Traits.IsAnonymousProxy}}\n            </td>\n          </tr>\n\n          <tr style=\"border: none;\">\n            <td style=\"border: none;\">\n              Satellite Provider\n            </td>\n            <td style=\"border: none;\">\n              {{.Traits.IsSatelliteProvider}}\n            </td>\n          </tr>\n        </table>  \n      </div>\n    </div>\n</div>\n{{end}}"
var _Assets2e41dd3189ac8c9e3dd205e726db965e20b2444c = "{{define \"base\"}}\n<!doctype html>\n<html lang=\"en\">\n  <head>\n    <!-- Required meta tags -->\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n\n    <!-- Bootstrap CSS -->\n    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css\" integrity=\"sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh\" crossorigin=\"anonymous\">\n\n    <title>GoIPInfo - {{template \"ipaddress\" .}}</title>\n  </head>\n  <body>\n    <nav class=\"navbar navbar-light bg-light\">\n      <form class=\"form-inline\">\n        <input class=\"form-control mr-sm-2\" type=\"search\" placeholder=\"IP Address\" aria-label=\"ip-address\">\n        <button class=\"btn btn-outline-success my-2 my-sm-0\" type=\"submit\">Get Info</button>\n      </form>\n    </nav>\n\n    {{template \"main\" .}}\n\n    <!-- Optional JavaScript -->\n    <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n    <script src=\"https://code.jquery.com/jquery-3.4.1.slim.min.js\" integrity=\"sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n\" crossorigin=\"anonymous\"></script>\n    <script src=\"https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js\" integrity=\"sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo\" crossorigin=\"anonymous\"></script>\n    <script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js\" integrity=\"sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6\" crossorigin=\"anonymous\"></script>\n  </body>\n</html>\n{{end}}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"ipAddressInfo.tmpl", "base.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1590186399, 1590186399186603904),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1590186182, 1590186182496614623),
		Data:     nil,
	}, "/templates/ipAddressInfo.tmpl": &assets.File{
		Path:     "/templates/ipAddressInfo.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1590186182, 1590186182496614623),
		Data:     []byte(_Assetscc9a4ea24b3abb724da85710e7e7309794800cc4),
	}, "/templates/base.tmpl": &assets.File{
		Path:     "/templates/base.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1590186182, 1590186182496614623),
		Data:     []byte(_Assets2e41dd3189ac8c9e3dd205e726db965e20b2444c),
	}}, "")