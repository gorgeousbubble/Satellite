using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace package
{
    public partial class FormMain : Form
    {
        public FormMain()
        {
            InitializeComponent();
        }

        private void FormMain_Load(object sender, EventArgs e)
        {
            // FormMain Initial
            AutoScaleMode = AutoScaleMode.Dpi;
            MaximizeBox = false;
            MinimizeBox = true;
            DoubleBuffered = true;
            StartPosition = FormStartPosition.CenterScreen;

            // TabPack Initial
            listView_pack.View = View.Details;
            listView_pack.Columns.Add("number", 60, HorizontalAlignment.Center);
            listView_pack.Columns.Add("files", 240, HorizontalAlignment.Center);
            listView_pack.Columns.Add("detials", 400, HorizontalAlignment.Center);

            comboBox_pack.DropDownStyle = ComboBoxStyle.DropDownList;
            comboBox_pack.Items.Clear();
            comboBox_pack.Items.Add("aes");
            comboBox_pack.Items.Add("3des");
            comboBox_pack.Items.Add("des");
            comboBox_pack.Items.Add("rsa");
            comboBox_pack.Items.Add("base64");
            comboBox_pack.SelectedIndex = 0;

        }

        private void FormMain_FormClosed(object sender, FormClosedEventArgs e)
        {

        }
    }
}
