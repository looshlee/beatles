/*
 *  Copyright (C) 2019 Authors of Cilium
 *
 *  This program is free software; you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation; either version 2 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program; if not, write to the Free Software
 *  Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 */
#include <bpf/ctx/skb.h>
#include <bpf/api.h>

#include <node_config.h>
#include <netdev_config.h>

#include "lib/common.h"

__section("to-host")
int to_host(struct __ctx_buff *ctx)
{
	// Upper 16 bits may carry proxy port number, clear it out
	__u32 magic = ctx->cb[0] & 0xFFFF;
	if (magic == MARK_MAGIC_TO_PROXY) {
		ctx->mark = ctx->cb[0];
		ctx->cb[0] = 0;
	}
	return CTX_ACT_OK;
}

BPF_LICENSE("GPL");
