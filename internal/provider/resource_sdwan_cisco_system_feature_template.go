// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &CiscoSystemFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CiscoSystemFeatureTemplateResource{}

func NewCiscoSystemFeatureTemplateResource() resource.Resource {
	return &CiscoSystemFeatureTemplateResource{}
}

type CiscoSystemFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CiscoSystemFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_system_feature_template"
}

func (r *CiscoSystemFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cisco System feature template.").AddMinimumVersionDescription("15.0.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Required:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of supported device types").AddStringEnumDescription("vedge-C8000V", "vedge-C8300-1N1S-4T2X", "vedge-C8300-1N1S-6T", "vedge-C8300-2N2S-6T", "vedge-C8300-2N2S-4T2X", "vedge-C8500-12X4QC", "vedge-C8500-12X", "vedge-C8500-20X6C", "vedge-C8500L-8S4X", "vedge-C8200-1N-4T", "vedge-C8200L-1N-4T").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the timezone").AddStringEnumDescription("Europe/Andorra", "Asia/Dubai", "Asia/Kabul", "America/Antigua", "America/Anguilla", "Europe/Tirane", "Asia/Yerevan", "Africa/Luanda", "Antarctica/McMurdo", "Antarctica/Rothera", "Antarctica/Palmer", "Antarctica/Mawson", "Antarctica/Davis", "Antarctica/Casey", "Antarctica/Vostok", "Antarctica/DumontDUrville", "Antarctica/Syowa", "America/Argentina/Buenos_Aires", "America/Argentina/Cordoba", "America/Argentina/Salta", "America/Argentina/Jujuy", "America/Argentina/Tucuman", "America/Argentina/Catamarca", "America/Argentina/La_Rioja", "America/Argentina/San_Juan", "America/Argentina/Mendoza", "America/Argentina/San_Luis", "America/Argentina/Rio_Gallegos", "America/Argentina/Ushuaia", "Pacific/Pago_Pago", "Europe/Vienna", "Australia/Lord_Howe", "Antarctica/Macquarie", "Australia/Hobart", "Australia/Currie", "Australia/Melbourne", "Australia/Sydney", "Australia/Broken_Hill", "Australia/Brisbane", "Australia/Lindeman", "Australia/Adelaide", "Australia/Darwin", "Australia/Perth", "Australia/Eucla", "America/Aruba", "Europe/Mariehamn", "Asia/Baku", "Europe/Sarajevo", "America/Barbados", "Asia/Dhaka", "Europe/Brussels", "Africa/Ouagadougou", "Europe/Sofia", "Asia/Bahrain", "Africa/Bujumbura", "Africa/Porto-Novo", "America/St_Barthelemy", "Atlantic/Bermuda", "Asia/Brunei", "America/La_Paz", "America/Kralendijk", "America/Noronha", "America/Belem", "America/Fortaleza", "America/Recife", "America/Araguaina", "America/Maceio", "America/Bahia", "America/Sao_Paulo", "America/Campo_Grande", "America/Cuiaba", "America/Santarem", "America/Porto_Velho", "America/Boa_Vista", "America/Manaus", "America/Eirunepe", "America/Rio_Branco", "America/Nassau", "Asia/Thimphu", "Africa/Gaborone", "Europe/Minsk", "America/Belize", "America/St_Johns", "America/Halifax", "America/Glace_Bay", "America/Moncton", "America/Goose_Bay", "America/Blanc-Sablon", "America/Toronto", "America/Nipigon", "America/Thunder_Bay", "America/Iqaluit", "America/Pangnirtung", "America/Resolute", "America/Atikokan", "America/Rankin_Inlet", "America/Winnipeg", "America/Rainy_River", "America/Regina", "America/Swift_Current", "America/Edmonton", "America/Cambridge_Bay", "America/Yellowknife", "America/Inuvik", "America/Creston", "America/Dawson_Creek", "America/Vancouver", "America/Whitehorse", "America/Dawson", "Indian/Cocos", "Africa/Kinshasa", "Africa/Lubumbashi", "Africa/Bangui", "Africa/Brazzaville", "Europe/Zurich", "Africa/Abidjan", "Pacific/Rarotonga", "America/Santiago", "Pacific/Easter", "Africa/Douala", "Asia/Shanghai", "Asia/Harbin", "Asia/Chongqing", "Asia/Urumqi", "Asia/Kashgar", "America/Bogota", "America/Costa_Rica", "America/Havana", "Atlantic/Cape_Verde", "America/Curacao", "Indian/Christmas", "Asia/Nicosia", "Europe/Prague", "Europe/Berlin", "Europe/Busingen", "Africa/Djibouti", "Europe/Copenhagen", "America/Dominica", "America/Santo_Domingo", "Africa/Algiers", "America/Guayaquil", "Pacific/Galapagos", "Europe/Tallinn", "Africa/Cairo", "Africa/El_Aaiun", "Africa/Asmara", "Europe/Madrid", "Africa/Ceuta", "Atlantic/Canary", "Africa/Addis_Ababa", "Europe/Helsinki", "Pacific/Fiji", "Atlantic/Stanley", "Pacific/Chuuk", "Pacific/Pohnpei", "Pacific/Kosrae", "Atlantic/Faroe", "Europe/Paris", "Africa/Libreville", "Europe/London", "America/Grenada", "Asia/Tbilisi", "America/Cayenne", "Europe/Guernsey", "Africa/Accra", "Europe/Gibraltar", "America/Godthab", "America/Danmarkshavn", "America/Scoresbysund", "America/Thule", "Africa/Banjul", "Africa/Conakry", "America/Guadeloupe", "Africa/Malabo", "Europe/Athens", "Atlantic/South_Georgia", "America/Guatemala", "Pacific/Guam", "Africa/Bissau", "America/Guyana", "Asia/Hong_Kong", "America/Tegucigalpa", "Europe/Zagreb", "America/Port-au-Prince", "Europe/Budapest", "Asia/Jakarta", "Asia/Pontianak", "Asia/Makassar", "Asia/Jayapura", "Europe/Dublin", "Asia/Jerusalem", "Europe/Isle_of_Man", "Asia/Kolkata", "Indian/Chagos", "Asia/Baghdad", "Asia/Tehran", "Atlantic/Reykjavik", "Europe/Rome", "Europe/Jersey", "America/Jamaica", "Asia/Amman", "Asia/Tokyo", "Africa/Nairobi", "Asia/Bishkek", "Asia/Phnom_Penh", "Pacific/Tarawa", "Pacific/Enderbury", "Pacific/Kiritimati", "Indian/Comoro", "America/St_Kitts", "Asia/Pyongyang", "Asia/Seoul", "Asia/Kuwait", "America/Cayman", "Asia/Almaty", "Asia/Qyzylorda", "Asia/Aqtobe", "Asia/Aqtau", "Asia/Oral", "Asia/Vientiane", "Asia/Beirut", "America/St_Lucia", "Europe/Vaduz", "Asia/Colombo", "Africa/Monrovia", "Africa/Maseru", "Europe/Vilnius", "Europe/Luxembourg", "Europe/Riga", "Africa/Tripoli", "Africa/Casablanca", "Europe/Monaco", "Europe/Chisinau", "Europe/Podgorica", "America/Marigot", "Indian/Antananarivo", "Pacific/Majuro", "Pacific/Kwajalein", "Europe/Skopje", "Africa/Bamako", "Asia/Rangoon", "Asia/Ulaanbaatar", "Asia/Hovd", "Asia/Choibalsan", "Asia/Macau", "Pacific/Saipan", "America/Martinique", "Africa/Nouakchott", "America/Montserrat", "Europe/Malta", "Indian/Mauritius", "Indian/Maldives", "Africa/Blantyre", "America/Mexico_City", "America/Cancun", "America/Merida", "America/Monterrey", "America/Matamoros", "America/Mazatlan", "America/Chihuahua", "America/Ojinaga", "America/Hermosillo", "America/Tijuana", "America/Santa_Isabel", "America/Bahia_Banderas", "Asia/Kuala_Lumpur", "Asia/Kuching", "Africa/Maputo", "Africa/Windhoek", "Pacific/Noumea", "Africa/Niamey", "Pacific/Norfolk", "Africa/Lagos", "America/Managua", "Europe/Amsterdam", "Europe/Oslo", "Asia/Kathmandu", "Pacific/Nauru", "Pacific/Niue", "Pacific/Auckland", "Pacific/Chatham", "Asia/Muscat", "America/Panama", "America/Lima", "Pacific/Tahiti", "Pacific/Marquesas", "Pacific/Gambier", "Pacific/Port_Moresby", "Asia/Manila", "Asia/Karachi", "Europe/Warsaw", "America/Miquelon", "Pacific/Pitcairn", "America/Puerto_Rico", "Asia/Gaza", "Asia/Hebron", "Europe/Lisbon", "Atlantic/Madeira", "Atlantic/Azores", "Pacific/Palau", "America/Asuncion", "Asia/Qatar", "Indian/Reunion", "Europe/Bucharest", "Europe/Belgrade", "Europe/Kaliningrad", "Europe/Moscow", "Europe/Volgograd", "Europe/Samara", "Asia/Yekaterinburg", "Asia/Omsk", "Asia/Novosibirsk", "Asia/Novokuznetsk", "Asia/Krasnoyarsk", "Asia/Irkutsk", "Asia/Yakutsk", "Asia/Khandyga", "Asia/Vladivostok", "Asia/Sakhalin", "Asia/Ust-Nera", "Asia/Magadan", "Asia/Kamchatka", "Asia/Anadyr", "Africa/Kigali", "Asia/Riyadh", "Pacific/Guadalcanal", "Indian/Mahe", "Africa/Khartoum", "Europe/Stockholm", "Asia/Singapore", "Atlantic/St_Helena", "Europe/Ljubljana", "Arctic/Longyearbyen", "Europe/Bratislava", "Africa/Freetown", "Europe/San_Marino", "Africa/Dakar", "Africa/Mogadishu", "America/Paramaribo", "Africa/Juba", "Africa/Sao_Tome", "America/El_Salvador", "America/Lower_Princes", "Asia/Damascus", "Africa/Mbabane", "America/Grand_Turk", "Africa/Ndjamena", "Indian/Kerguelen", "Africa/Lome", "Asia/Bangkok", "Asia/Dushanbe", "Pacific/Fakaofo", "Asia/Dili", "Asia/Ashgabat", "Africa/Tunis", "Pacific/Tongatapu", "Europe/Istanbul", "America/Port_of_Spain", "Pacific/Funafuti", "Asia/Taipei", "Africa/Dar_es_Salaam", "Europe/Kiev", "Europe/Uzhgorod", "Europe/Zaporozhye", "Europe/Simferopol", "Africa/Kampala", "Pacific/Johnston", "Pacific/Midway", "Pacific/Wake", "America/New_York", "America/Detroit", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Indiana/Indianapolis", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Vevay", "America/Chicago", "America/Indiana/Tell_City", "America/Indiana/Knox", "America/Menominee", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/North_Dakota/Beulah", "America/Denver", "America/Boise", "America/Phoenix", "America/Los_Angeles", "America/Anchorage", "America/Juneau", "America/Sitka", "America/Yakutat", "America/Nome", "America/Adak", "America/Metlakatla", "Pacific/Honolulu", "America/Montevideo", "Asia/Samarkand", "Asia/Tashkent", "Europe/Vatican", "America/St_Vincent", "America/Caracas", "America/Tortola", "America/St_Thomas", "Asia/Ho_Chi_Minh", "Pacific/Efate", "Pacific/Wallis", "Pacific/Apia", "Asia/Aden", "Indian/Mayotte", "Africa/Johannesburg", "Africa/Lusaka", "Africa/Harare", "UTC").AddDefaultValueDescription("UTC").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Europe/Andorra", "Asia/Dubai", "Asia/Kabul", "America/Antigua", "America/Anguilla", "Europe/Tirane", "Asia/Yerevan", "Africa/Luanda", "Antarctica/McMurdo", "Antarctica/Rothera", "Antarctica/Palmer", "Antarctica/Mawson", "Antarctica/Davis", "Antarctica/Casey", "Antarctica/Vostok", "Antarctica/DumontDUrville", "Antarctica/Syowa", "America/Argentina/Buenos_Aires", "America/Argentina/Cordoba", "America/Argentina/Salta", "America/Argentina/Jujuy", "America/Argentina/Tucuman", "America/Argentina/Catamarca", "America/Argentina/La_Rioja", "America/Argentina/San_Juan", "America/Argentina/Mendoza", "America/Argentina/San_Luis", "America/Argentina/Rio_Gallegos", "America/Argentina/Ushuaia", "Pacific/Pago_Pago", "Europe/Vienna", "Australia/Lord_Howe", "Antarctica/Macquarie", "Australia/Hobart", "Australia/Currie", "Australia/Melbourne", "Australia/Sydney", "Australia/Broken_Hill", "Australia/Brisbane", "Australia/Lindeman", "Australia/Adelaide", "Australia/Darwin", "Australia/Perth", "Australia/Eucla", "America/Aruba", "Europe/Mariehamn", "Asia/Baku", "Europe/Sarajevo", "America/Barbados", "Asia/Dhaka", "Europe/Brussels", "Africa/Ouagadougou", "Europe/Sofia", "Asia/Bahrain", "Africa/Bujumbura", "Africa/Porto-Novo", "America/St_Barthelemy", "Atlantic/Bermuda", "Asia/Brunei", "America/La_Paz", "America/Kralendijk", "America/Noronha", "America/Belem", "America/Fortaleza", "America/Recife", "America/Araguaina", "America/Maceio", "America/Bahia", "America/Sao_Paulo", "America/Campo_Grande", "America/Cuiaba", "America/Santarem", "America/Porto_Velho", "America/Boa_Vista", "America/Manaus", "America/Eirunepe", "America/Rio_Branco", "America/Nassau", "Asia/Thimphu", "Africa/Gaborone", "Europe/Minsk", "America/Belize", "America/St_Johns", "America/Halifax", "America/Glace_Bay", "America/Moncton", "America/Goose_Bay", "America/Blanc-Sablon", "America/Toronto", "America/Nipigon", "America/Thunder_Bay", "America/Iqaluit", "America/Pangnirtung", "America/Resolute", "America/Atikokan", "America/Rankin_Inlet", "America/Winnipeg", "America/Rainy_River", "America/Regina", "America/Swift_Current", "America/Edmonton", "America/Cambridge_Bay", "America/Yellowknife", "America/Inuvik", "America/Creston", "America/Dawson_Creek", "America/Vancouver", "America/Whitehorse", "America/Dawson", "Indian/Cocos", "Africa/Kinshasa", "Africa/Lubumbashi", "Africa/Bangui", "Africa/Brazzaville", "Europe/Zurich", "Africa/Abidjan", "Pacific/Rarotonga", "America/Santiago", "Pacific/Easter", "Africa/Douala", "Asia/Shanghai", "Asia/Harbin", "Asia/Chongqing", "Asia/Urumqi", "Asia/Kashgar", "America/Bogota", "America/Costa_Rica", "America/Havana", "Atlantic/Cape_Verde", "America/Curacao", "Indian/Christmas", "Asia/Nicosia", "Europe/Prague", "Europe/Berlin", "Europe/Busingen", "Africa/Djibouti", "Europe/Copenhagen", "America/Dominica", "America/Santo_Domingo", "Africa/Algiers", "America/Guayaquil", "Pacific/Galapagos", "Europe/Tallinn", "Africa/Cairo", "Africa/El_Aaiun", "Africa/Asmara", "Europe/Madrid", "Africa/Ceuta", "Atlantic/Canary", "Africa/Addis_Ababa", "Europe/Helsinki", "Pacific/Fiji", "Atlantic/Stanley", "Pacific/Chuuk", "Pacific/Pohnpei", "Pacific/Kosrae", "Atlantic/Faroe", "Europe/Paris", "Africa/Libreville", "Europe/London", "America/Grenada", "Asia/Tbilisi", "America/Cayenne", "Europe/Guernsey", "Africa/Accra", "Europe/Gibraltar", "America/Godthab", "America/Danmarkshavn", "America/Scoresbysund", "America/Thule", "Africa/Banjul", "Africa/Conakry", "America/Guadeloupe", "Africa/Malabo", "Europe/Athens", "Atlantic/South_Georgia", "America/Guatemala", "Pacific/Guam", "Africa/Bissau", "America/Guyana", "Asia/Hong_Kong", "America/Tegucigalpa", "Europe/Zagreb", "America/Port-au-Prince", "Europe/Budapest", "Asia/Jakarta", "Asia/Pontianak", "Asia/Makassar", "Asia/Jayapura", "Europe/Dublin", "Asia/Jerusalem", "Europe/Isle_of_Man", "Asia/Kolkata", "Indian/Chagos", "Asia/Baghdad", "Asia/Tehran", "Atlantic/Reykjavik", "Europe/Rome", "Europe/Jersey", "America/Jamaica", "Asia/Amman", "Asia/Tokyo", "Africa/Nairobi", "Asia/Bishkek", "Asia/Phnom_Penh", "Pacific/Tarawa", "Pacific/Enderbury", "Pacific/Kiritimati", "Indian/Comoro", "America/St_Kitts", "Asia/Pyongyang", "Asia/Seoul", "Asia/Kuwait", "America/Cayman", "Asia/Almaty", "Asia/Qyzylorda", "Asia/Aqtobe", "Asia/Aqtau", "Asia/Oral", "Asia/Vientiane", "Asia/Beirut", "America/St_Lucia", "Europe/Vaduz", "Asia/Colombo", "Africa/Monrovia", "Africa/Maseru", "Europe/Vilnius", "Europe/Luxembourg", "Europe/Riga", "Africa/Tripoli", "Africa/Casablanca", "Europe/Monaco", "Europe/Chisinau", "Europe/Podgorica", "America/Marigot", "Indian/Antananarivo", "Pacific/Majuro", "Pacific/Kwajalein", "Europe/Skopje", "Africa/Bamako", "Asia/Rangoon", "Asia/Ulaanbaatar", "Asia/Hovd", "Asia/Choibalsan", "Asia/Macau", "Pacific/Saipan", "America/Martinique", "Africa/Nouakchott", "America/Montserrat", "Europe/Malta", "Indian/Mauritius", "Indian/Maldives", "Africa/Blantyre", "America/Mexico_City", "America/Cancun", "America/Merida", "America/Monterrey", "America/Matamoros", "America/Mazatlan", "America/Chihuahua", "America/Ojinaga", "America/Hermosillo", "America/Tijuana", "America/Santa_Isabel", "America/Bahia_Banderas", "Asia/Kuala_Lumpur", "Asia/Kuching", "Africa/Maputo", "Africa/Windhoek", "Pacific/Noumea", "Africa/Niamey", "Pacific/Norfolk", "Africa/Lagos", "America/Managua", "Europe/Amsterdam", "Europe/Oslo", "Asia/Kathmandu", "Pacific/Nauru", "Pacific/Niue", "Pacific/Auckland", "Pacific/Chatham", "Asia/Muscat", "America/Panama", "America/Lima", "Pacific/Tahiti", "Pacific/Marquesas", "Pacific/Gambier", "Pacific/Port_Moresby", "Asia/Manila", "Asia/Karachi", "Europe/Warsaw", "America/Miquelon", "Pacific/Pitcairn", "America/Puerto_Rico", "Asia/Gaza", "Asia/Hebron", "Europe/Lisbon", "Atlantic/Madeira", "Atlantic/Azores", "Pacific/Palau", "America/Asuncion", "Asia/Qatar", "Indian/Reunion", "Europe/Bucharest", "Europe/Belgrade", "Europe/Kaliningrad", "Europe/Moscow", "Europe/Volgograd", "Europe/Samara", "Asia/Yekaterinburg", "Asia/Omsk", "Asia/Novosibirsk", "Asia/Novokuznetsk", "Asia/Krasnoyarsk", "Asia/Irkutsk", "Asia/Yakutsk", "Asia/Khandyga", "Asia/Vladivostok", "Asia/Sakhalin", "Asia/Ust-Nera", "Asia/Magadan", "Asia/Kamchatka", "Asia/Anadyr", "Africa/Kigali", "Asia/Riyadh", "Pacific/Guadalcanal", "Indian/Mahe", "Africa/Khartoum", "Europe/Stockholm", "Asia/Singapore", "Atlantic/St_Helena", "Europe/Ljubljana", "Arctic/Longyearbyen", "Europe/Bratislava", "Africa/Freetown", "Europe/San_Marino", "Africa/Dakar", "Africa/Mogadishu", "America/Paramaribo", "Africa/Juba", "Africa/Sao_Tome", "America/El_Salvador", "America/Lower_Princes", "Asia/Damascus", "Africa/Mbabane", "America/Grand_Turk", "Africa/Ndjamena", "Indian/Kerguelen", "Africa/Lome", "Asia/Bangkok", "Asia/Dushanbe", "Pacific/Fakaofo", "Asia/Dili", "Asia/Ashgabat", "Africa/Tunis", "Pacific/Tongatapu", "Europe/Istanbul", "America/Port_of_Spain", "Pacific/Funafuti", "Asia/Taipei", "Africa/Dar_es_Salaam", "Europe/Kiev", "Europe/Uzhgorod", "Europe/Zaporozhye", "Europe/Simferopol", "Africa/Kampala", "Pacific/Johnston", "Pacific/Midway", "Pacific/Wake", "America/New_York", "America/Detroit", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Indiana/Indianapolis", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Vevay", "America/Chicago", "America/Indiana/Tell_City", "America/Indiana/Knox", "America/Menominee", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/North_Dakota/Beulah", "America/Denver", "America/Boise", "America/Phoenix", "America/Los_Angeles", "America/Anchorage", "America/Juneau", "America/Sitka", "America/Yakutat", "America/Nome", "America/Adak", "America/Metlakatla", "Pacific/Honolulu", "America/Montevideo", "Asia/Samarkand", "Asia/Tashkent", "Europe/Vatican", "America/St_Vincent", "America/Caracas", "America/Tortola", "America/St_Thomas", "Asia/Ho_Chi_Minh", "Pacific/Efate", "Pacific/Wallis", "Pacific/Apia", "Asia/Aden", "Indian/Mayotte", "Africa/Johannesburg", "Africa/Lusaka", "Africa/Harare", "UTC"),
				},
			},
			"timezone_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the hostname").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"hostname_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"system_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set a text description of the device").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"system_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the location of the device").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
			},
			"location_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"latitude": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the device’s physical latitude").AddFloatRangeDescription(-90, 90).String,
				Optional:            true,
				Validators: []validator.Float64{
					float64validator.Between(-90, 90),
				},
			},
			"latitude_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"longitude": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the device’s physical longitude").AddFloatRangeDescription(-180, 180).String,
				Optional:            true,
				Validators: []validator.Float64{
					float64validator.Between(-180, 180),
				},
			},
			"longitude_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"geo_fencing": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Geo fencing").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"geo_fencing_range": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the device’s geo fencing range").AddIntegerRangeDescription(100, 10000).AddDefaultValueDescription("100").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(100, 10000),
				},
			},
			"geo_fencing_range_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"geo_fencing_sms": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Geo fencing").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"geo_fencing_sms_phone_numbers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set device’s geo fencing SMS phone number").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"number": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Mobile number, ex: +1231234414").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(7, 20),
							},
						},
						"number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"device_groups": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device groups (Use comma(,) for multiple groups)").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"device_groups_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"controller_group_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure a list of comma-separated device groups").String,
				ElementType:         types.Int64Type,
				Optional:            true,
			},
			"controller_group_list_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"system_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the system IP address").String,
				Optional:            true,
			},
			"system_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"overlay_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the Overlay ID").AddIntegerRangeDescription(1, 4294967295).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"overlay_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"site_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the site identifier").AddIntegerRangeDescription(1, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"site_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"port_offset": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the TLOC port offset when multiple devices are behind a NAT").AddIntegerRangeDescription(0, 19).AddDefaultValueDescription("0").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 19),
				},
			},
			"port_offset_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"port_hopping": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable port hopping").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"port_hopping_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"control_session_pps": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the policer rate for control sessions").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("300").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"control_session_pps_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"track_transport": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure tracking of transport").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"track_transport_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"track_interface_tag": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("OMP Tag attached to routes based on interface tracking").AddIntegerRangeDescription(1, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"track_interface_tag_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"console_baud_rate": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the console baud rate").AddStringEnumDescription("1200", "2400", "4800", "9600", "19200", "38400", "57600", "115200").AddDefaultValueDescription("9600").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1200", "2400", "4800", "9600", "19200", "38400", "57600", "115200"),
				},
			},
			"console_baud_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"max_omp_sessions": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the maximum number of OMP sessions <1..100> the device can have").AddIntegerRangeDescription(1, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
			},
			"max_omp_sessions_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"multi_tenant": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device is multi-tenant").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"multi_tenant_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"track_default_gateway": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable default gateway tracking").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"track_default_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"admin_tech_on_failure": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Collect admin-tech before reboot due to daemon failure").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"admin_tech_on_failure_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Idle CLI timeout in minutes").AddIntegerRangeDescription(0, 300).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 300),
				},
			},
			"idle_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"trackers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tracker configuration").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tracker name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
							},
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"endpoint_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address of endpoint").String,
							Optional:            true,
						},
						"endpoint_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"endpoint_dns_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("DNS name of endpoint").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 63),
							},
						},
						"endpoint_dns_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"endpoint_api_url": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("API url of endpoint").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 512),
							},
						},
						"endpoint_api_url_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"elements": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tracker member names separated by space").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"elements_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"boolean": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of grouping to be performed for tracker group").AddStringEnumDescription("or", "and").AddDefaultValueDescription("or").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("or", "and"),
							},
						},
						"boolean_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"threshold": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Probe Timeout threshold <100..1000> milliseconds").AddIntegerRangeDescription(100, 1000).AddDefaultValueDescription("300").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(100, 1000),
							},
						},
						"threshold_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Probe interval <10..600> seconds").AddIntegerRangeDescription(20, 600).AddDefaultValueDescription("60").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(20, 600),
							},
						},
						"interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"multiplier": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Probe failure multiplier <1..10> failed attempts").AddIntegerRangeDescription(1, 10).AddDefaultValueDescription("3").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 10),
							},
						},
						"multiplier_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Default(Interface)").AddStringEnumDescription("tracker-group", "interface", "static-route").AddDefaultValueDescription("interface").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("tracker-group", "interface", "static-route"),
							},
						},
						"type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"object_trackers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Object Track configuration").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"object_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Object tracker ID").AddIntegerRangeDescription(1, 1000).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 1000),
							},
						},
						"object_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("interface name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"sig": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("service sig").String,
							Optional:            true,
						},
						"sig_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address of route").String,
							Optional:            true,
						},
						"ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"mask": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Route Ip Mask").AddDefaultValueDescription("0.0.0.0").String,
							Optional:            true,
						},
						"mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VPN").AddIntegerRangeDescription(0, 65527).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65527),
							},
						},
						"group_tracks_ids": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tracks id in group configuration").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"track_id": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Track id").AddIntegerRangeDescription(1, 1000).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 1000),
										},
									},
									"track_id_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"boolean": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of grouping to be performed for tracker group").AddStringEnumDescription("and", "or").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("and", "or"),
							},
						},
						"boolean_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"on_demand_tunnel": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable On-demand Tunnel").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"on_demand_tunnel_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"on_demand_tunnel_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Idle CLI timeout in minutes").AddIntegerRangeDescription(0, 300).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 300),
				},
			},
			"on_demand_tunnel_idle_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"region_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set region ID").AddIntegerRangeDescription(1, 63).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 63),
				},
			},
			"region_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"secondary_region_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set secondary region ID").AddIntegerRangeDescription(1, 63).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 63),
				},
			},
			"secondary_region_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"role": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the role for router").AddStringEnumDescription("edge-router", "border-router").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("edge-router", "border-router"),
				},
			},
			"role_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"affinity_group_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the affinity group number for router").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"affinity_group_number_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"affinity_group_preference": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the affinity group preference").String,
				ElementType:         types.Int64Type,
				Optional:            true,
			},
			"affinity_group_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"transport_gateway": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable transport gateway").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"transport_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_mrf_migration": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable migration mode to Multi-Region Fabric").AddStringEnumDescription("enabled", "enabled-from-bgp-core").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "enabled-from-bgp-core"),
				},
			},
			"migration_bgp_community": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set BGP community during migration from BGP-core based network").AddIntegerRangeDescription(1, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
		},
	}
}

func (r *CiscoSystemFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CiscoSystemFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiscoSystem

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/feature", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("templateId").String())
	plan.Version = types.Int64Value(0)
	plan.TemplateType = types.StringValue(plan.getModel())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *CiscoSystemFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiscoSystem

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/feature/object/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid Template Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *CiscoSystemFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CiscoSystem

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	body := plan.toBody(ctx)
	r.updateMutex.Lock()
	res, err := r.client.Put("/template/feature/"+url.QueryEscape(plan.Id.ValueString()), body)
	r.updateMutex.Unlock()
	if err != nil {
		if res.Get("error.message").String() == "Template locked in edit mode." {
			resp.Diagnostics.AddWarning("Client Warning", fmt.Sprintf("Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again."))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	if plan.hasChanges(ctx, &state) {
		plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)
	} else {
		plan.Version = state.Version
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *CiscoSystemFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiscoSystem

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/feature/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *CiscoSystemFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
