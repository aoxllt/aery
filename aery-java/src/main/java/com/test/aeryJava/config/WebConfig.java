package com.test.aeryJava.config;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.CorsRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class WebConfig implements WebMvcConfigurer {

    @Override
    public void addCorsMappings(CorsRegistry registry) {
        registry.addMapping("/**")
                .allowedOriginPatterns("*")  // 允许所有域名
                .allowedMethods("GET", "POST", "PUT", "DELETE")
                .allowCredentials(true)  // 允许发送凭证
                .maxAge(3600);  // 预检请求缓存时间
    }
}