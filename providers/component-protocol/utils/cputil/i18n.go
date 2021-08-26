// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cputil

import (
	"context"

	"github.com/erda-project/erda-infra/providers/i18n"
)

// Language .
func Language(ctx context.Context) i18n.LanguageCodes {
	lang := GetHeader(ctx, "Lang")
	if len(lang) <= 0 {
		lang = GetHeader(ctx, "Accept-Language")
	}
	langs, _ := i18n.ParseLanguageCode(lang)
	return langs
}

// I18n .
func I18n(ctx context.Context, key string, args ...interface{}) string {
	return SDK(ctx).I18n(key, args...)
}
