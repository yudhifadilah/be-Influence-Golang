-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 15, 2025 at 01:46 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `starpowers`
--

-- --------------------------------------------------------

--
-- Table structure for table `articles`
--

CREATE TABLE `articles` (
  `id` int(6) UNSIGNED NOT NULL,
  `title` varchar(255) NOT NULL,
  `excerpt` text NOT NULL,
  `content` text NOT NULL,
  `image` varchar(255) NOT NULL,
  `reg_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `articles`
--

INSERT INTO `articles` (`id`, `title`, `excerpt`, `content`, `image`, `reg_date`) VALUES
(1, 'coba', 'coba', 'coba', 'uploads/download (4).jpg', '2025-02-05 17:15:43'),
(2, 'tes', 'tes', 'tes', 'uploads/download (3).jpg', '2025-02-05 17:18:38'),
(3, 'tes1', 'tes1', 'tes1', 'uploads/download (3).jpg', '2025-02-05 17:21:36'),
(4, 'qwe', 'qwe', 'qwe', 'uploads/download.jpg', '2025-02-05 17:24:11'),
(5, 'Makin Berpengaruh! Tren Influencer Marketing yang Mendominasi 2025', 'Influencer marketing terus menjadi strategi utama dalam dunia pemasaran', 'contoh ', 'uploads/3fd6c3a6d7cc73953be7cbaff47e4b95d276a0a6_xxl-1.jpg', '2025-02-07 09:18:24');

-- --------------------------------------------------------

--
-- Table structure for table `bank_accounts`
--

CREATE TABLE `bank_accounts` (
  `id` int(11) NOT NULL,
  `bank_type` varchar(255) NOT NULL,
  `account_number` varchar(255) NOT NULL,
  `account_holder` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `bank_accounts`
--

INSERT INTO `bank_accounts` (`id`, `bank_type`, `account_number`, `account_holder`) VALUES
(2, 'Mandiri', '233402', 'FADHIL');

-- --------------------------------------------------------

--
-- Table structure for table `brands`
--

CREATE TABLE `brands` (
  `id` int(11) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role` varchar(255) NOT NULL,
  `brand_name` varchar(255) DEFAULT NULL,
  `pic_name` varchar(255) DEFAULT NULL,
  `pic_phone` varchar(20) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `referral_code` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `brand_logo` varchar(255) DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `brands`
--

INSERT INTO `brands` (`id`, `email`, `password`, `role`, `brand_name`, `pic_name`, `pic_phone`, `province`, `city`, `referral_code`, `created_at`, `name`, `phone`, `address`, `brand_logo`, `updated_at`, `deleted_at`) VALUES
(1, 'hayangwonton@gmail.com', '$2y$10$SOMcGz8PYt0gngNQIrYBTedMAznrE8CELgN2qAlvE3qAYidkUALtW', '', 'tes', 'tes', 'tes', '16', '213', '', '2025-02-05 18:10:53', '', NULL, '', NULL, NULL, NULL),
(2, 'niki@gmail.com', '$2y$10$oLV6MBabIH.VuQKpe2E7KOfzt4uZefL7ZT4WUFn4F.9YnCUuo/CX2', '', 'niki', 'niki', '0823232323', '9', '64', '', '2025-02-07 08:38:22', '', NULL, '', NULL, NULL, NULL),
(3, 'shifa@gmail.com', '$2y$10$r2XBRk8Vng.PJXjb5F/CM.bf6pQyARhDKojxA7ZYKQusGflZ6tXkS', '', 'Shifa', 'Fadhil', '0822232323', '9', '64', '', '2025-02-07 09:22:42', '', NULL, '', NULL, NULL, NULL),
(4, 'laksana@gmail.com', '$2y$10$Xbwz7qL0rFyt0l.3np4coeyQFaqv6Cp4bhUbLPmxkZqYIQ/leZf5W', '', 'laksana', 'laksana', '0987654', '9', '73', '', '2025-02-07 12:09:22', '', NULL, '', NULL, NULL, NULL),
(5, '123333test@example.com', '$2y$10$sRtc4VvXsLFs8zVb0lks3eHjcTqc7BGtrmR7I2GzsV1bG9gPJUR..', '', 'john', 'eeeeee', '08123456789', 'jambi', 'bandung', NULL, '2025-02-14 12:38:25', NULL, NULL, NULL, 'http://localhost:8080/uploads/brands/1739561905_32d835252b8c58461498.png', '2025-02-14 12:38:25', NULL),
(6, '', '$2a$10$Ll09sEtopE6IQa0n5Jk2YeMT5uwoxgzzgiY17yuNMWQrXZYOAi/P2', '', '', '', '', '', '', NULL, '2025-02-15 10:44:24', NULL, NULL, NULL, 'uploads/brands/LOGO LADIES.png', '2025-02-15 10:44:24', NULL),
(7, '', '$2a$10$yKieyaZGBW/vivu8AfvRD.pnCRlaoN8NibePM8pAqEPLMDgCqE4B.', 'brand', '', '', '', '', '', NULL, '2025-02-15 12:38:16', NULL, NULL, NULL, 'uploads/brands/LOGO LADIES.png', '2025-02-15 12:38:16', NULL),
(8, 'tatang@gmail.com', '$2a$10$UjSMG2Hg./PukfDH.m9Q6.XVhVCiSdR5ghqCTv2ZgNG.AEnzpLGXi', 'brand', 'john', 'eeeeee', '08123456789', 'jambi', 'bandung', NULL, '2025-02-15 12:40:47', NULL, NULL, NULL, 'uploads/brands/LOGO LADIES.png', '2025-02-15 12:40:47', NULL),
(9, 'tatang@gmail.com', '$2a$10$49QW4ufy7Vcd0SoLL80Fe.RtaMtJJERq8a9XDkHNnLCJ5.YMN7J/m', 'brand', 'john', 'eeeeee', '08123456789', 'jambi', 'bandung', NULL, '2025-02-15 12:41:49', NULL, NULL, NULL, 'uploads/brands/LOGO LADIES.png', '2025-02-15 12:41:49', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `influencer_id` int(11) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `pdf_file` varchar(255) NOT NULL,
  `status` enum('pending','accepted','rejected') DEFAULT 'pending',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `campaigns`
--

INSERT INTO `campaigns` (`id`, `name`, `category`, `influencer_id`, `start_date`, `end_date`, `pdf_file`, `status`, `created_at`, `updated_at`) VALUES
(1, '123333test@example.com', 'Entertainment', 2, '2025-03-01', '2025-03-10', 'http://localhost:8080/uploads/pdf/1739562971_89bb2020e07a36a9aa1a.pdf', 'pending', '2025-02-14 19:56:11', '2025-02-14 19:56:11'),
(2, '123333test@example.com', 'enterntainment', 5, '2025-03-01', '2025-03-10', 'http://localhost:8080/uploads/pdf/1739563241_ac99bb654fe32e62999f.pdf', 'pending', '2025-02-14 20:00:41', '2025-02-14 20:00:41'),
(3, 'wwweqewqeqweqw', 'enterntainment', 5, '2025-03-01', '2025-03-10', '', '', '2025-02-15 12:01:46', '2025-02-15 12:01:46'),
(4, 'wwweqewqeqweqw', 'enterntainment', 5, '2025-03-01', '2025-03-10', '', '', '2025-02-15 12:04:25', '2025-02-15 12:04:25'),
(5, 'wwweqewqeqweqw', 'enterntainment', 5, '2025-03-01', '2025-03-10', 'uploads\\pdf\\Proposal-OfficialRoomCard_Starzone.pdf', '', '2025-02-15 12:05:02', '2025-02-15 12:05:02'),
(6, 'wwweqewqeqweqw', 'enterntainment', 5, '2025-03-01', '2025-03-10', 'uploads\\pdf\\5_Proposal-OfficialRoomCard_Starzone.pdf', 'pending', '2025-02-15 12:07:39', '2025-02-15 12:07:39');

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`) VALUES
(1, 'Entertainment'),
(2, 'Lifestyle and Travel'),
(3, 'Family and Parenting'),
(4, 'Beauty and Fashion'),
(5, 'Health and Support'),
(6, 'Technology'),
(7, 'Food and Beverages'),
(8, 'Gaming');

-- --------------------------------------------------------

--
-- Table structure for table `faqs`
--

CREATE TABLE `faqs` (
  `id` int(11) NOT NULL,
  `category` varchar(50) NOT NULL,
  `question` text NOT NULL,
  `answer` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `faqs`
--

INSERT INTO `faqs` (`id`, `category`, `question`, `answer`) VALUES
(2, 'influencer', 'Apa itu Starpowers Indonesia?', 'Starpowers merupakan platform influencer marketing  yang mempertemukan antara influencer dan brand seluruh Indonesia dalam suatu campaign digita'),
(3, 'influencer', 'Apa saja keuntungan bergabung di Starpowers ?', 'Banyak keuntungan yang bisa kamu dapatkan jika bergabung di Starpowers, di antaranya: Kesempatan berkolaborasi dengan brand seluruh Indonesia. Sebagai sarana menampilkan portofolio endorsement. Jaminan pembayaran fee kamu, karena sistem pembayaran yang transparan.'),
(4, 'influencer', 'Berapa biaya untuk yang dibutuhkan untuk bergabung di Starpowers ?', 'Kamu tidak dipungut biaya sama sekali saat bergabung di Starpowers.'),
(5, 'influencer', 'Apa saja syarat untuk bergabung di Starpowers ?', 'Tidak ada syarat khusus untuk bergabung di Starpowers. Namun pastikan memiliki akun media sosial yang aktif dan rekening bank untuk mempermudah penarikan dana saat campaign telah berhasil.'),
(6, 'influencer', 'Bagaimana cara bergabung menjadi influencer Starpowers ?', 'Kamu bisa akses website Starpowers, Lalu pilih \"Daftar Sebagai Influencer\" dan ikuti alur registrasinya.'),
(7, 'influencer', 'Bagaimana cara menyambungkan akun media sosial di Starpowers ?', 'Kamu dapat menyambungkan akun media sosial di Starpowers dengan mengikuti langkah-langkah berikut ini: Buka halaman Akun pada Website Starpowers. Cek bagian Aset dan pilih akun media sosial yang ingin kamu sambungkan. Masukkan Username pada akun media sosial yang ingin kamu sambungkan. Pilih Izinkan untuk menghubungkan akun kamu. Jika sudah berhasil tersambung, isi nominal harga tiap post yang sesuai dengan rate card kamu. Simpan perubahan kamu dengan klik Save.'),
(8, 'influencer', 'Bagaimana cara mengubah akun media sosial di Starpowers ?', 'Kamu dapat mengganti akun media sosial di Starpowers dengan mengikuti langkah-langkah berikut ini: Buka halaman Akun pada Website Starpowers. Pilih akun media sosial yang ingin kamu ubah/hapus, klik Edit. Pilih bagian Putus Aset. Jika sudah terputus, sambungkan kembali dengan akun media sosial kamu yang baru. Pilih Izinkan untuk menghubungkan akun kamu. Simpan perubahan kamu dengan klik Save.'),
(9, 'influencer', 'Bagaimana cara mengikuti campaign di Starpowers ?', 'Cara mengikuti campaign sangatlah mudah. Kamu bisa ikuti langkah-langkah berikut ini: Login akun kamu ke Website Starpowers. Pilih campaign yang ingin kamu ikuti di tab Campaign. Baca dan cermati detail dan brief campaign. Jika sudah sesuai dengan keinginan kamu, klik Ikuti Campaign. Submit proposal konten dan contoh caption yang akan kamu post. Tunggu approval dari brand sebelum konten diupload. Jika konten telah di-approve, kamu bisa upload kontennya sesuai brief. Jangan lupa untuk kirim screenshot bukti upload konten dan tunggu hingga campaign berakhir.'),
(10, 'influencer', 'Bagaimana cara membatalkan campaign yang saya ikuti ?', 'Untuk membatalkan campaign, kamu dapat mengabaikan campaign tersebut hingga masa berlakunya berakhir. Sehingga kamu akan dianggap tidak mengikuti campaign.'),
(11, 'influencer', 'Kapan dana masuk ke rekening saya ?', 'Dana yang sudah ditarik (withdraw) akan diproses ke rekening kamu tiap Jumat. Pastikan kamu mengajukan withdraw sebelum Jumat pukul 12.00, karena jika lebih dari itu maka transfer akan diproses di minggu depan. Perhatikan bahwa: Jumlah min. dana yang dapat ditarik adalah Rp100.000. Dana yang ditarik belum termasuk potongan admin fee sebesar 10% dan Pph 21 sebesar 3%. Jika dana ditransfer ke bank non BCA, akan dikenai biaya transfer sebesar Rp5.000. Dana akan masuk ke rekening paling lambat 10 hari kerja setelah pengajuan withdraw tergantung kebijakan masing-masing bank. Pastikan data rekeningmu sudah benar agar meminimalisir kesalahan dalam proses transfer.'),
(12, 'brand', 'Apa itu Starpowers Marketplace ? ', 'Starpowers Marketplace merupakan platform influencer marketing pertama di Indonesia yang mempertemukan antara influencer dan brand seluruh Indonesia dalam suatu campaign digital.'),
(13, 'brand', 'Apa saja keuntungan bergabung di Starpowers ?', 'Banyak keuntungan yang bisa kamu dapatkan jika bergabung di Allstars, antaranya: Kemudahan akses pada platform influencer marketing. Dapat terhubung langsung dengan influencer. Jaminan keamanan saldo kamu, karena jika influencer tidak melakukan campaign sesuai brief maka fee tidak akan dibayarkan'),
(14, 'brand', 'Berapa biaya untuk yang dibutuhkan untuk bergabung di Starpowers  ?', 'Kamu tidak dipungut biaya sama sekali saat bergabung di Starpowers.'),
(15, 'brand', 'Bagaimana cara mendaftarkan brand ke Starpowers ? ', 'Kamu bisa mendownload Website Starpowers. Lalu pilih \"Daftar Sebagai Brand\" dan ikuti alur pendaftarannya'),
(16, 'brand', 'Bagaimana cara membuat campaign di Starpowers ?', 'Cara membuat campaign sangatlah mudah. Kamu dapat ikuti langkah-langkah berikut ini: Login akun kamu ke Website Starpowers. Klik tab Buat Campaign. Lengkapi informasi Nama Campaign, Nama Brand, dan Kategori Brand. Pilih jenis campaign yang sesuai Lengkapi informasi Kriteria Influencer, Jumlah & Harga Post, dan Brief Campaign. Jika sudah sesuai dengan keinginan kamu, klik Konfirmasi Campaign. Akan muncul rincian biaya campaign dan jika sudah sesuai dengan keinginanmu, klik Proses Pembayaran. Pilih metode pembayaran dan segera bayar dengan nominal yang sudah ditentukan. Jika pembayaran sudah selesai, campaign dapat dijalankan.'),
(17, 'brand', 'Apa saya bisa me-refund sisa saldo saya ?   ', 'Saldo yang sudah masuk di Starpowers tidak dapat ditarik kembali, namun saldo tersebut dapat digunakan untuk pembuatan campaign baru dengan hanya menambah jumlah nominal yang kurang');

-- --------------------------------------------------------

--
-- Table structure for table `influencers`
--

CREATE TABLE `influencers` (
  `id` int(11) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `birth_date` date NOT NULL,
  `gender` varchar(50) NOT NULL,
  `influencer_category` varchar(255) NOT NULL,
  `phone_number` varchar(20) NOT NULL,
  `referral_code` varchar(50) DEFAULT NULL,
  `ktp_number` varchar(50) NOT NULL,
  `npwp_number` varchar(50) NOT NULL,
  `instagram_link` varchar(255) NOT NULL,
  `followers_count` int(11) NOT NULL,
  `profile_picture` varchar(255) DEFAULT NULL,
  `bank_account` varchar(255) DEFAULT NULL,
  `account_number` varchar(255) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `registration_date` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `influencers`
--

INSERT INTO `influencers` (`id`, `email`, `password`, `role`, `full_name`, `birth_date`, `gender`, `influencer_category`, `phone_number`, `referral_code`, `ktp_number`, `npwp_number`, `instagram_link`, `followers_count`, `profile_picture`, `bank_account`, `account_number`, `province`, `city`, `registration_date`, `updated_at`) VALUES
(1, 'fadhilirmani10@gmail.com', '$2y$10$MmciHVIyzdkVjfG6tVI0heqr96YKpYkePnj3l7iP67r9nVed21TFO', '', 'fadhil', '2025-01-31', 'male', 'Entertainment', '123', '', '123', '123', '123', 119, 'uploads/download (4).jpg', '123', '123', '9', '63', '2025-02-05 18:13:13', NULL),
(2, 'shifa1@gmail.com', '$2y$10$GmNP6YINpL7T1WgWHdsM7ODdi0nslE3QUywpWlCND0Gh6uTsuJKvi', '', 'laksana', '2025-01-28', 'female', 'Entertainment', '123456', '', '45678', '12344', 'https://www.instagram.com/', 10000, 'uploads/IMG_2132.JPG', 'bca', '234678', '18', '236', '2025-02-07 11:46:21', NULL),
(3, 'test@example.com', '$2y$10$hAJVmeq3CXy0HeLneRMl/O.JXyKOIn9PVPQCXwBKqS/uzYV6daPzm', '', 'John Doe', '1995-01-15', 'male', 'Tech', '08123456789', NULL, '1234567890123456', '987654321098765', 'https://instagram.com/johndoe', 10000, NULL, NULL, NULL, NULL, NULL, '2025-02-14 11:12:29', '2025-02-14 18:12:29'),
(4, '123test@example.com', '$2y$10$bKzeBbdAxb7.n/UghaWXnesQiJXA/OiNM1CrHxph/LeKLawAMLhg.', '', 'john', '1995-01-15', 'female', 'enterntainment', '08123456789', NULL, '1234567890123456', '987654321098765', 'https://instagram.com/johndoe', 100000, 'http://localhost:8080/uploads/1739557292_c352f2d35fb960277763.jpg', NULL, NULL, NULL, NULL, '2025-02-14 11:21:32', '2025-02-14 18:21:32'),
(5, '123333test@example.com', '$2y$10$971MUjFOYUCJEm6Xn4JIeOwDYSsH5ubroE2M1Dut3pfG.p7EGlxxO', '', 'john', '1995-01-15', 'female', 'enterntainment', '08123456789', NULL, '1234567890123456', '987654321098765', 'https://instagram.com/johndoe', 100000, 'http://localhost:8080/uploads/1739557405_ff357da24860fb3ba70f.jpg', NULL, NULL, NULL, NULL, '2025-02-14 11:23:25', '2025-02-14 18:23:25'),
(6, '12333eeeeee3test@example.com', '$2y$10$PlHFI5jsrHqP3ZQrInPe7O5eps5sir8kqS/v65TzQH73hhdDeA9oy', '', 'john', '1995-01-15', 'female', 'enterntainment', '08123456789', NULL, '1234567890123456', '987654321098765', 'https://instagram.com/johndoe', 100000, 'http://localhost:8080/uploads/1739608613_3728742682e21d72bffe.jpg', NULL, NULL, NULL, NULL, '2025-02-15 01:36:53', '2025-02-15 08:36:53'),
(7, '', '$2a$10$GAkU2dMDQHdtpQv.Ki5FAevXz2rM7wAwNEqeO6Ug5O1CtftH2jFey', '', '', '0000-00-00', '', '', '', NULL, '', '', '', 0, 'uploads\\wp2695919-black-lagoon-wallpaper-1080p.jpg', '', '', '', '', '2025-02-15 10:53:52', NULL),
(8, '', '$2a$10$T1RbwbEOSmXe1JX/U7SdH.Rw5m9rft6MMjD6vTudQm1CbcrFql0Bu', '', '', '0000-00-00', '', '', '', NULL, '', '', '', 0, 'uploads\\wp2695919-black-lagoon-wallpaper-1080p.jpg', '', '', '', '', '2025-02-15 11:50:38', NULL),
(9, 'fayudhi@gmail.com', '$2a$10$xBgzEOcv1T/PYEeVW3687ulWN2F/JjY7eTdNB4h7voc16KW2MsBZ.', '', 'john', '1995-01-15', 'female', 'enterntainment', '08123456789', NULL, '1234567890123456', '987654321098765', 'https://instagram.com/johndoe', 100000, 'uploads\\wp2695919-black-lagoon-wallpaper-1080p.jpg', '', '', '', '', '2025-02-15 11:58:07', NULL),
(10, 'dapit@gmail.com', '$2a$10$JVDc8BFfVXW1oZ3rS0g0qujHDZb5CxDIHpXEAWFehiuUPn1u.Y08e', 'influencer', 'john', '1995-01-15', 'female', 'enterntainment', '08123456789', NULL, '1234567890123456', '987654321098765', 'https://instagram.com/johndoe', 100000, 'uploads\\wp2695919-black-lagoon-wallpaper-1080p.jpg', '', '', '', '', '2025-02-15 12:31:11', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `services`
--

CREATE TABLE `services` (
  `id` int(11) NOT NULL,
  `influencer_id` int(11) NOT NULL,
  `service_name` varchar(255) NOT NULL,
  `price_per_post` decimal(10,2) NOT NULL,
  `description` text NOT NULL,
  `duration` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `services`
--

INSERT INTO `services` (`id`, `influencer_id`, `service_name`, `price_per_post`, `description`, `duration`) VALUES
(1, 1, 'Promosi di Instagram Story', 200000.00, 'Promosi produk atau jasa di Instagram Story selama 24 jam.', 25);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `role`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'admin', '$2y$10$uryfmgscvCZYUxFx57pyKe.vKjoWWz0ItbjzIb0B.z3GRrPw6me8G', '', '2025-02-14 17:40:48', NULL, NULL),
(2, 'ewe', '$2y$10$F5vXOi./NVq8xLy6F3GljOVYUbOGwEz5n5owLy6OzXofQeQeiKO92', '', '2025-02-15 08:54:22', NULL, NULL),
(3, 'testuser', '$2a$10$13E7eP5wLdids0zQmP4dqeJCIQ1GeXyX58X7/EbaDr.ZrDdzNYqhS', 'admin', '2025-02-15 10:13:07', '2025-02-15 17:13:07', NULL),
(4, 'eeeeeeeee', '$2a$10$ihbUQNzE594rGr2bQ39pye5x.vTWptbx/MYvXEVUSwV2rT7gtJpGC', 'admin', '2025-02-15 10:13:26', '2025-02-15 17:13:26', NULL),
(5, 'eeeee222eeee', '$2a$10$6fxlB1E/.asSlEri0NAEWOYPyuI8uZAX9qdxPhIo3rMsRgtboNtX.', 'admin', '2025-02-15 10:39:51', '2025-02-15 17:39:51', NULL),
(6, 'dapi', '$2a$10$NoCPP/vdKYtDhnyk.Bjmj.g0Qeh7pA.qSAH0QnJJpY8v0VoRBK5rS', 'admin', '2025-02-15 10:56:40', '2025-02-15 17:56:40', NULL),
(7, 'dapit', '$2a$10$vXZu.hV19OaAWyrxm5aVuu8sZ054h6WcsMLSztFloWSaSkGEjU7ae', 'admin', '2025-02-15 10:56:46', '2025-02-15 17:56:46', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `bank_accounts`
--
ALTER TABLE `bank_accounts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `brands`
--
ALTER TABLE `brands`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`),
  ADD KEY `influencer_id` (`influencer_id`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `faqs`
--
ALTER TABLE `faqs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `influencers`
--
ALTER TABLE `influencers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `services`
--
ALTER TABLE `services`
  ADD PRIMARY KEY (`id`),
  ADD KEY `influencer_id` (`influencer_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `articles`
--
ALTER TABLE `articles`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `bank_accounts`
--
ALTER TABLE `bank_accounts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `brands`
--
ALTER TABLE `brands`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `faqs`
--
ALTER TABLE `faqs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `influencers`
--
ALTER TABLE `influencers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `services`
--
ALTER TABLE `services`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD CONSTRAINT `campaigns_ibfk_1` FOREIGN KEY (`influencer_id`) REFERENCES `influencers` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
