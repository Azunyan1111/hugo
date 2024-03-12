// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugolib

import (
	"sync"

	"github.com/Azunyan1111/hugo/common/maps"
	"github.com/Azunyan1111/hugo/compare"
	"github.com/Azunyan1111/hugo/lazy"
	"github.com/Azunyan1111/hugo/markup/converter"
	"github.com/Azunyan1111/hugo/navigation"
	"github.com/Azunyan1111/hugo/output/layouts"
	"github.com/Azunyan1111/hugo/resources/page"
	"github.com/Azunyan1111/hugo/resources/resource"
	"github.com/Azunyan1111/hugo/source"
)

type nextPrevProvider interface {
	getNextPrev() *nextPrev
}

func (p *pageCommon) getNextPrev() *nextPrev {
	return p.posNextPrev
}

type nextPrevInSectionProvider interface {
	getNextPrevInSection() *nextPrev
}

func (p *pageCommon) getNextPrevInSection() *nextPrev {
	return p.posNextPrevSection
}

type pageCommon struct {
	s *Site
	m *pageMeta

	sWrapped page.Site

	// Lazily initialized dependencies.
	init *lazy.Init

	// Store holds state that survives server rebuilds.
	store *maps.Scratch

	// All of these represents the common parts of a page.Page
	maps.Scratcher
	navigation.PageMenusProvider
	page.AuthorProvider
	page.AlternativeOutputFormatsProvider
	page.ChildCareProvider
	page.FileProvider
	page.GetPageProvider
	page.GitInfoProvider
	page.InSectionPositioner
	page.OutputFormatsProvider
	page.PageMetaProvider
	page.Positioner
	page.RawContentProvider
	page.RelatedKeywordsProvider
	page.RefProvider
	page.ShortcodeInfoProvider
	page.SitesProvider
	page.TranslationsProvider
	page.TreeProvider
	resource.LanguageProvider
	resource.ResourceDataProvider
	resource.ResourceNameTitleProvider
	resource.ResourceParamsProvider
	resource.ResourceTypeProvider
	resource.MediaTypeProvider
	resource.TranslationKeyProvider
	compare.Eqer

	// Describes how paths and URLs for this page and its descendants
	// should look like.
	targetPathDescriptor page.TargetPathDescriptor

	layoutDescriptor     layouts.LayoutDescriptor
	layoutDescriptorInit sync.Once

	// Set if feature enabled and this is in a Git repo.
	gitInfo    source.GitInfo
	codeowners []string

	// Positional navigation
	posNextPrev        *nextPrev
	posNextPrevSection *nextPrev

	// Menus
	pageMenus *pageMenus

	// Internal use
	page.InternalDependencies

	contentConverterInit sync.Once
	contentConverter     converter.Converter
}

func (p *pageCommon) Store() *maps.Scratch {
	return p.store
}
