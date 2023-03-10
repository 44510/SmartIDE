/*
SmartIDE - CLI
Copyright (C) 2023 leansoftX.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package compose

// 网络配置
type Network struct {
	// TODO driver
	// TODO driver_opts
	// TODO attachable
	// TODO enable_ipv6
	// TODO ipam
	// TODO internal
	// TODO labels
	External bool `yaml:"external,omitempty"` // 是否是外部创建
	// TODO name
}
