// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/damongolding/immich-kiosk/config"
	"github.com/damongolding/immich-kiosk/immich"
	"github.com/damongolding/immich-kiosk/utils"
)

// ImageLocation generates a formatted string of the image location based on EXIF information.
// It combines the city, state, and country information if available.
func ImageLocation(info immich.ExifInfo, hideCountries []string) string {
	var parts []string

	if info.City != "" {
		parts = append(parts, info.City)
	}

	if info.State != "" {
		parts = append(parts, info.State)
	}

	if info.Country != "" && !slices.Contains(hideCountries, strings.ToLower(info.Country)) {
		// Insert the line break if there are already parts (city or state)
		if len(parts) > 0 {
			parts = append(parts, "<span>, </span><br class=\"responsive-break\"/>"+info.Country)
		} else {
			parts = append(parts, info.Country)
		}
	}

	return strings.Join(parts, ", ")
}

// ImageExif generates a formatted string of EXIF information for an image.
// It includes f-number, exposure time, focal length, and ISO if available.
func ImageExif(info immich.ExifInfo) string {
	var stats strings.Builder

	if info.FNumber != 0 {
		stats.WriteString(fmt.Sprintf("<span class=\"image--metadata--exif--fnumber\">&#402;</span>/%.1f", info.FNumber))
	}

	if info.ExposureTime != "" {
		if stats.Len() > 0 {
			stats.WriteString("<span class=\"image--metadata--exif--seperator\">&#124;</span>")
		}
		stats.WriteString(fmt.Sprintf("%s<small>s</small>", info.ExposureTime))
	}

	if info.FocalLength != 0 {
		if stats.Len() > 0 {
			stats.WriteString("<span class=\"image--metadata--exif--seperator\">&#124;</span>")
		}
		stats.WriteString(fmt.Sprintf("%vmm", info.FocalLength))
	}

	if info.Iso != 0 {
		if stats.Len() > 0 {
			stats.WriteString("<span class=\"image--metadata--exif--seperator\">&#124;</span>")
		}
		stats.WriteString(fmt.Sprintf("ISO %v", info.Iso))
	}

	return stats.String()
}

// ImageDateTime generates a formatted date and time string for an image based on the view data settings.
// It can display date, time, or both, in various formats.
func ImageDateTime(viewData ViewData, imageIndex int) string {
	var imageDate string

	imageTimeFormat := "15:04"
	if viewData.ImageTimeFormat == "12" {
		imageTimeFormat = time.Kitchen
	}

	imageDateFormat := utils.DateToLayout(viewData.ImageDateFormat)
	if imageDateFormat == "" {
		imageDateFormat = config.DefaultDateLayout
	}

	switch {
	case (viewData.ShowImageDate && viewData.ShowImageTime):
		imageDate = fmt.Sprintf("%s %s", viewData.Images[imageIndex].ImmichImage.LocalDateTime.Format(imageTimeFormat), viewData.Images[imageIndex].ImmichImage.LocalDateTime.Format(imageDateFormat))
	case viewData.ShowImageDate:
		imageDate = fmt.Sprintf("%s", viewData.Images[imageIndex].ImmichImage.LocalDateTime.Format(imageDateFormat))
	case viewData.ShowImageTime:
		imageDate = fmt.Sprintf("%s", viewData.Images[imageIndex].ImmichImage.LocalDateTime.Format(imageTimeFormat))
	}

	return imageDate
}

// imageMetadata renders the metadata for an image, including date, time, EXIF information, location, and ID.
// The display of each piece of information is controlled by the ViewData settings.
func imageMetadata(viewData ViewData, imageIndex int) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{"image--metadata", fmt.Sprintf("image--metadata--theme-%s", viewData.Theme)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/views_image-metadata.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if viewData.ShowImageDate || viewData.ShowImageTime {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"image--metadata--date\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(ImageDateTime(viewData, imageIndex))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/views_image-metadata.templ`, Line: 105, Col: 41}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if viewData.ShowImageExif {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"image--metadata--exif\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.Raw(ImageExif(viewData.Images[imageIndex].ImmichImage.ExifInfo)).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if viewData.ShowImageLocation {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"image--metadata--location\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.Raw(ImageLocation(viewData.Images[imageIndex].ImmichImage.ExifInfo, viewData.HideCountries)).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if viewData.ShowImageID {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"image--metadata--id\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(viewData.Images[imageIndex].ImmichImage.ID)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/views_image-metadata.templ`, Line: 120, Col: 48}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
