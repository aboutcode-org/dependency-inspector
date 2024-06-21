===================
DepLock
===================

|license| |version| |build| 

.. |license| image:: https://img.shields.io/badge/License-Apache--2.0-blue.svg?style=for-the-badge
    :target: https://opensource.org/licenses/Apache-2.0

.. |version| image:: https://img.shields.io/github/v/release/nexB/dependency-inspector?style=for-the-badge

.. |build| image:: https://img.shields.io/github/actions/workflow/status/nexB/dependency-inspector/ci.yml?style=for-the-badge&logo=github

**DepLock** (Dependency Locker) is a command line tool to generate package lockfiles, 
and parsable package manifests to make it possible and easy to collect resolved dependencies
and accurate metadata for a project. It uses the standard package management tool for
each package type or ecosystem.

This tool is designed for use before a scan and is deployed as a single,
standalone executable file. It should be running in a shell or command prompt
configured to run the project build, with all required package management tools
pre-installed and configured.

The main usage is to run as a front command before a ScanCode.io or ScanCode Toolkit scan
but it is generic and can be used with any other tool. 

Supported Ecosystems
=====================

- **npm**: https://www.npmjs.com/
- **pnpm**: https://pnpm.io/
- **yarn**: https://yarnpkg.com/
- **swift**: https://www.swift.org/documentation/package-manager/
- **cocoapods**: https://cocoapods.org/
- **pypi**: https://pypi.org/
- **nuget**: https://www.nuget.org/


Installation
============

To install DepLock, follow these steps:

.. code-block:: bash

    # Download the latest binary depending on your OS and architecture
    # Replace <OS> with your operating system (linux, darwin, windows)
    # Replace <ARCH> with your architecture (amd64, arm64)
    curl -LO https://github.com/nexB/dependency-inspector/releases/latest/download/deplock-<OS>-<ARCH>

    # Make the binary executable
    chmod +x deplock-<OS>-<ARCH>

    # [Optional] Rename and move the binary to your PATH
    mv deplock-<OS>-<ARCH> /usr/local/bin/deplock

Alternatively you can build it from source as exaplained below.


Requirements
============

To use DepLock, you need to ensure that you have installed and configured
the package manager(s) your project needs.

For instance, `npm` must be installed to generate `npm` lockfiles.


Usage
=======

Here is how to get started and use the various commands:

.. code-block:: bash

    # Display the general help for DepLock
    deplock --help

    # Display help for a specific command
    deplock [command] --help

Example
-------

Generating lockfile for an npm project:

.. code-block:: bash

    # Generate lockfile in the current directory
    deplock npm

    # Generate lockfile in specified directory
    deplock npm /path/to/project

    # Forcefully generate lockfile, ignoring existing lockfiles
    deplock npm /path/to/project --force


Contribution
=============

We welcome contributions from the community! If you find a bug or have an idea for a new feature, 
please open an issue on the GitHub repository. If you want to contribute code, you can fork the
repository, make your changes, and submit a pull request.

- Please try to write a good commit message, see `good commit message wiki.
  <https://aboutcode.readthedocs.io/en/latest/contributing/writing_good_commit_messages.html>`_
- Add DCO Sign Off to your commits.

Development setup
------------------
Run these commands, starting from a git clone of https://github.com/nexB/dependency-inspector.git

.. code-block:: bash

    make dev

- Compile and run:

  .. code-block:: bash

     $ go run main.go

- Create binary:

  .. code-block:: bash

     $ make build

- Run tests:

  .. code-block:: bash

     $ make test


License
=======

SPDX-License-Identifier: Apache-2.0

DepLock is licensed under Apache License version 2.0.

.. code-block:: none

    You may not use this software except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
