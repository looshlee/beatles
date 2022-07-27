/*
 *  Copyright (C) 2016-2020 Authors of Cilium
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
#ifndef __BPF_API__
#define __BPF_API__

#include <linux/types.h>
#include <linux/byteorder.h>
#include <linux/bpf.h>
#include <linux/if_packet.h>

#include "compiler.h"
#include "section.h"
#include "helpers.h"
#include "builtins.h"

#define PIN_NONE		0
#define PIN_OBJECT_NS		1
#define PIN_GLOBAL_NS		2

struct bpf_elf_map {
	__u32 type;
	__u32 size_key;
	__u32 size_value;
	__u32 max_elem;
	__u32 flags;
	__u32 id;
	__u32 pinning;
#ifdef SOCKMAP
	__u32 inner_id;
	__u32 inner_idx;
#endif
};

#endif /* __BPF_API__ */
