components {
  id: "bullet"
  component: "/main/bullet.script"
}
components {
  id: "explosion"
  component: "/main/explosion.particlefx"
  position {
    x: -1.0
    y: -3.0
  }
}
components {
  id: "boom"
  component: "/main/boom.sound"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bullet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"bullets\"\n"
  "mask: \"enemies\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -1.0\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 3.760961\n"
  "  data: 9.010936\n"
  "  data: 10.32\n"
  "}\n"
  ""
}
