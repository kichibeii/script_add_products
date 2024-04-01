CREATE TABLE `products` (
  `id` bigint(20) NOT NULL,
  `title` varchar(255) NOT NULL,
  `shopify_id` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `product_images` (
  `id` bigint(20) NOT NULL,
  `product_id` bigint(2) NOT NULL,
  `link` varchar(500) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `product_variants` (
  `id` bigint(20) NOT NULL,
  `product_id` bigint(2) NOT NULL,
  `title` varchar(255) NOT NULL,
  `price` bigint(20) NOT NULL,
  `quantity` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
