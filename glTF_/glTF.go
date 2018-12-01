package glTF_

import (
	"encoding/json"
	"errors"

	"./accessors"
	"./animations"
	"./asset"
	"./bufferViews"
	"./buffers"
	"./images"
	"./materials"
	"./meshes"
	"./nodes"
	"./samplers"
	"./scenes"
	"./skins"
	"./textures"
)

type GLTF struct {
	//ok
	Accessors   []accessors.T
	Animations  []animations.T
	Asset       asset.T
	BufferViews []bufferViews.T
	Buffers     []buffers.T
	Images      []images.T
	Materials   []materials.T
	Meshes      []meshes.T
	Nodes       []nodes.T
	Samplers    []samplers.T
	Scene       int
	Scenes      []scenes.T
	Skins       []skins.T
	Textures    []textures.T
}

func (this *GLTF) init() {
}

func (this *GLTF) Compiler(gltfBytes []byte) error {
	var r interface{}

	json.Unmarshal(gltfBytes, &r)

	loli, ok := r.(map[string]interface{})

	if ok {
		for k, v := range loli {
			switch k {
			case "accessors":
				{
					for _, iv := range v.([]interface{}) {
						var temp accessors.T
						for j, jv := range iv.(map[string]interface{}) {
							switch j {
							case "bufferView":
								{
									temp.BufferView = int(jv.(float64))
									break
								}
							case "byteOffset":
								{
									temp.ByteOffset = int(jv.(float64))
									break
								}
							case "componentType":
								{
									temp.ComponentType = int(jv.(float64))
									break
								}
							case "count":
								{
									temp.Count = int(jv.(float64))
									break
								}
							case "max":
								{
									for _, lv := range jv.([]interface{}) {
										temp.Max = append(temp.Max, lv.(float64))
									}
									break
								}
							case "min":
								{
									for _, lv := range jv.([]interface{}) {
										temp.Min = append(temp.Min, lv.(float64))
									}
									break
								}
							case "type":
								{
									temp.Type = jv.(string)
									break
								}
							case "sparse":
								{
									for l, lv := range jv.(map[string]interface{}) {
										switch l {
										case "count":
											{
												temp.Sparse.Count = int(lv.(float64))
												break
											}
										case "indices":
											{
												for m, mv := range lv.(map[string]interface{}) {
													switch m {
													case "bufferView":
														{
															temp.Sparse.Indices.BufferView = int(mv.(float64))
															break
														}
													case "byteOffset":
														{
															temp.Sparse.Indices.ByteOffset = int(mv.(float64))
															break
														}
													case "componentType":
														{
															temp.Sparse.Indices.ComponentType = int(mv.(float64))
															break
														}
													}
												}
												break
											}
										case "values":
											{
												for m, mv := range lv.(map[string]interface{}) {
													switch m {
													case "bufferView":
														{
															temp.Sparse.Values.BufferView = int(mv.(float64))
															break
														}
													case "byteOffset":
														{
															temp.Sparse.Values.ByteOffset = int(mv.(float64))
															break
														}
													}
												}
												break
											}
										}
									}
									break
								}
							}
						}
						this.Accessors = append(this.Accessors, temp)
					}
					break
				}
			case "animations":
				{
					for _, iv := range v.([]interface{}) {
						var temp animations.T
						for j, jv := range iv.(map[string]interface{}) {
							switch j {
							case "channels":
								{

									for _, lv := range jv.([]interface{}) {
										var temp_c animations.T_channels
										for m, mv := range lv.(map[string]interface{}) {
											switch m {
											case "sampler":
												{
													temp_c.Sampler = int(mv.(float64))
													break
												}
											case "target":
												{
													for n, nv := range mv.(map[string]interface{}) {
														switch n {
														case "node":
															{
																temp_c.Target.Node = int(nv.(float64))
																break
															}
														case "path":
															{
																temp_c.Target.Path = nv.(string)
																break
															}
														}
													}
													break
												}
											}
										}
										temp.Channels = append(temp.Channels, temp_c)
									}
									break
								}
							case "name":
								{
									temp.Name = jv.(string)
								}
							case "samplers":
								{

									for _, lv := range jv.([]interface{}) {
										var temp_s animations.T_samplers
										for m, mv := range lv.(map[string]interface{}) {
											switch m {
											case "inpur":
												{
													temp_s.Input = int(mv.(float64))
													break
												}
											case "interpolation":
												{
													temp_s.Interpolation = mv.(string)
													break
												}
											case "Output":
												{
													temp_s.Output = int(mv.(float64))
													break
												}
											}
										}
										temp.Samplers = append(temp.Samplers, temp_s)
									}

									break
								}
							}
						}
						this.Animations = append(this.Animations, temp)
					}
				}

			}

		}

		return nil
	} else {
		var err error = errors.New("glTF decoding failure")
		return err
	}
}
