// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package course_of_action

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type CourseOfActionType struct {
	common.CommonObjectPropertiesType
	common.DescriptivePropertiesType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() CourseOfActionType {
	var obj CourseOfActionType
	obj.InitNewObject("course-of-action")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - CourseOfActionType
// ----------------------------------------------------------------------
