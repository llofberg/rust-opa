// main.rs

extern "C" {
    pub fn compile(p0: *const u8, p1: *const u8);
    pub fn query(p0: *const u8, p1: *const u8) -> u8;
}

fn main() {
    let module = b"\n\
      package example\n\
      default allow = false\n\
      allow {\n\
        input.identity = \"admin\"\n\
      }\n\
      allow {\n\
        input.method = \"GET\"\n\
      }\n\0";

    let f = b"example.rego\0";
    let q = b"data.example.allow\0";

    let input1 = b"{\"identity\": \"bob\", \"method\":   \"GET\"}\0";
    let input2 = b"{\"identity\": \"bob\", \"method\":   \"POST\"}\0";

    unsafe {
        compile(module.as_ptr(), f.as_ptr());
    };

    let r = unsafe { query(q.as_ptr(), input1.as_ptr()) };
    println!("input1: {}", r);

    let r = unsafe { query(q.as_ptr(), input2.as_ptr()) };
    println!("input2: {}", r);
}
