// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Interface for all executors (Java/Python/Go/SCIO)
package executors

type executor interface {
	// Validate validates executable file.
	// Return result of validation (true/false) and error if it occurs
	Validate(filePath string) (bool, error)

	// Compile compiles executable file.
	// Return error if it occurs
	Compile(filePath string) error

	// Run runs executable file.
	// Return logs and error if it occurs
	Run(filePath string) (string, error)
}
