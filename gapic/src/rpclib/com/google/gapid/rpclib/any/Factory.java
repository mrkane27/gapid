/*
 * Copyright (C) 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package com.google.gapid.rpclib.any;

public final class Factory {
    public static void register() {
        //<<<Start:Java.FactoryBody:2>>>
        BoolSlice.register();
        Bool.register();
        Float32Slice.register();
        Float32.register();
        Float64Slice.register();
        Float64.register();
        IdSlice.register();
        Id.register();
        Int16Slice.register();
        Int16.register();
        Int32Slice.register();
        Int32.register();
        Int64Slice.register();
        Int64.register();
        Int8Slice.register();
        Int8.register();
        ObjectSlice.register();
        ObjectBox.register();
        StringSlice.register();
        StringBox.register();
        Uint16Slice.register();
        Uint16.register();
        Uint32Slice.register();
        Uint32.register();
        Uint64Slice.register();
        Uint64.register();
        Uint8Slice.register();
        Uint8.register();
        //<<<End:Java.FactoryBody:2>>>
    }
}
